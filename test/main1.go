package main

import (
	"encoding/json"
	"fmt"
)

// main.go

type Row struct {
	Id       int
	ParentId int
	Name     string
}

type Node struct {
	*Row
	Children []*Node `json:"children,omitempty"`
}

func generateTree(rows []*Row) []*Node {
	idMap := make(map[int]*Node, len(rows)+1)
	topNode := &Node{Row: &Row{Id: 0, ParentId: 0, Name: "root"}}
	idMap[topNode.Id] = topNode

	for _, row := range rows {
		// 先判断 当前ID 对应的节点是否存在
		nNode, ok := idMap[row.Id]
		if ok {
			// 存在就补录节点数据
			nNode.Row = row
		} else {
			// 不存在就新建一个节点，并放入map
			nNode = &Node{Row: row}
			idMap[nNode.Id] = nNode
		}
		// 根据parentId 获取 parentNode
		pNode, ok := idMap[nNode.ParentId]
		if !ok { // 获取不到，说明父节点还没有，先进行添加
			pNode = &Node{}
			idMap[nNode.ParentId] = pNode
		}

		// 将当前节点添加到父节点的 array 中
		pNode.Children = append(pNode.Children, nNode)
	}
	return topNode.Children
}

func main2() {
	data := []*Row{
		{Id: 2, ParentId: 1, Name: "java开发"},
		{Id: 7, ParentId: 3, Name: "python书"},
		{Id: 3, ParentId: 1, Name: "python开发"},
		{Id: 5, ParentId: 4, Name: "水果销售"},
		{Id: 6, ParentId: 4, Name: "蔬菜销售"},
		{Id: 1, ParentId: 0, Name: "开发菜单"},
		{Id: 4, ParentId: 0, Name: "销售菜单"},
	}

	tree := generateTree(data)
	marshal, err := json.MarshalIndent(tree, "", "  ")
	//marshal, err := json.Marshal(tree)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(marshal))

	fmt.Printf("%s\n", marshal)

}
