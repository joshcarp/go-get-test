package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"github.com/stretchr/testify/require"
)

type CategoryNode struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	ParentID string `json:"parentId"`
}

type CategoryMap map[string]CategoryNode

type CategoryStore struct {
	AssignmentMap map[string][]string
	HierarchyMap  map[string]map[string]CategoryNode
}

/*

category keys [1 3 6 9]
hierarchyMap keys [ANZx HEM]

category keys [3 6 7 9]
hierarchyMap keys [ANZx HEM]
*/

func buildCategoryMaps(categories CategoryMap) (map[string][]string, map[string]map[string]CategoryNode) {
	assignmentMap := make(map[string][]string)
	hierarchyMap := make(map[string]map[string]CategoryNode)
	leafNodes := make(CategoryMap)
	for k, v := range categories {
		leafNodes[k] = v
	}
	// Build HierarchyMap
	//for _, categoryID := range SortedKeys(categories) {
	//	categoryNode := categories[categoryID]
	for categoryID, categoryNode := range categories {
		//categoryNode := categories[categoryID]
		// Group CategoryNodes by their type
		_, initialised := hierarchyMap[categoryNode.Type]
		if !initialised {
			hierarchyMap[categoryNode.Type] = make(map[string]CategoryNode)
		}
		hierarchyMap[categoryNode.Type][categoryID] = categoryNode

		// Remove parents from leafNodes
		delete(leafNodes, categoryNode.ParentID)
	}
	fmt.Println("category keys", SortedKeys(categories))
	fmt.Println("hierarchyMap keys", SortedKeys(hierarchyMap))

	// Build Assignment Map
	for categoryID, categoryNode := range leafNodes {
		assignmentMap[categoryNode.Type] = append(assignmentMap[categoryNode.Type], categoryID)
	}

	return assignmentMap, hierarchyMap
}

func TestBuildCategoryMaps(t *testing.T) {
	tests := []struct {
		Name          string
		CategoryMap   CategoryMap
		AssignmentMap map[string][]string
		HierarchyMap  map[string]map[string]CategoryNode
	}{
		{
			Name: "Returns expected results",
			CategoryMap: CategoryMap{
				"1": {Type: "ANZx", Name: "One", ParentID: ""},
				"2": {Type: "ANZx", Name: "Two", ParentID: "1"},
				"3": {Type: "ANZx", Name: "Three", ParentID: "2"},
				"4": {Type: "ANZx", Name: "Four", ParentID: ""},
				"5": {Type: "ANZx", Name: "Five", ParentID: "4"},
				"6": {Type: "ANZx", Name: "Six", ParentID: "5"},
				"7": {Type: "HEM", Name: "Seven", ParentID: ""},
				"8": {Type: "HEM", Name: "Eight", ParentID: "7"},
				"9": {Type: "HEM", Name: "Nine", ParentID: "8"},
			},
			AssignmentMap: map[string][]string{
				"ANZx": {"3", "6"},
				"HEM":  {"9"},
			},
			HierarchyMap: map[string]map[string]CategoryNode{
				"ANZx": {
					"1": {Type: "ANZx", Name: "One", ParentID: ""},
					"2": {Type: "ANZx", Name: "Two", ParentID: "1"},
					"3": {Type: "ANZx", Name: "Three", ParentID: "2"},
					"4": {Type: "ANZx", Name: "Four", ParentID: ""},
					"5": {Type: "ANZx", Name: "Five", ParentID: "4"},
					"6": {Type: "ANZx", Name: "Six", ParentID: "5"},
				},
				"HEM": {
					"7": {Type: "HEM", Name: "Seven", ParentID: ""},
					"8": {Type: "HEM", Name: "Eight", ParentID: "7"},
					"9": {Type: "HEM", Name: "Nine", ParentID: "8"},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			aMap, hMap := buildCategoryMaps(tt.CategoryMap)
			require.EqualValues(t, tt.AssignmentMap, aMap)
			require.EqualValues(t, tt.HierarchyMap, hMap)
		})
	}
}

func SortedKeys(m interface{}) []string {
	keys := reflect.ValueOf(m).MapKeys()
	ret := make([]string, 0, len(keys))
	for _, v := range keys {
		ret = append(ret, v.String())
	}
	sort.Strings(ret)
	return ret
}