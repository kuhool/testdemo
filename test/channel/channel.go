package channel

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Name           string  `json:"name,omitempty"`
	Type           string  `json:"type,omitempty"`
	NodeID         string  `json:"nodeId,omitempty"`
	PrevID         string  `json:"prevId,omitempty"`
	ChildNode      *Node   `json:"childNode,omitempty"`
	ConditionNodes []*Node `json:"conditionNodes,omitempty"`
}

func ChannelTest() error {
	nodes := make([]*Node, 0)
	str := `[{"nodeId":"开始","type":"starter","approver":"328e4658-1446-4493-b9e6-97a76e4be092","aproverType":"","memberCount":0,"level":0,"actType":"","actioners":null},{"nodeId":"QVgaH92","type":"target_manager","approver":"14de967b-e7a6-4326-be6e-93e65ff8bc2b","aproverType":"approver","memberCount":1,"level":0,"actType":"and","actioners":[{"UserId":"14de967b-e7a6-4326-be6e-93e65ff8bc2b","userName":"admin_luo"}]},{"nodeId":"结束","type":"","approver":"","aproverType":"","memberCount":0,"level":0,"actType":"","actioners":null}]`
	err := json.Unmarshal([]byte(str), &nodes)
	if err != nil {
		return err
	}
	fmt.Printf("nodes节点: %v\n", nodes)
	getNodesChan := func() <-chan *Node {
		nodesChan := make(chan *Node, len(nodes))
		go func() {
			defer close(nodesChan)
			for _, v := range nodes {
				nodesChan <- v
			}
		}()
		return nodesChan
	}

	nodesChan := getNodesChan()
	for node := range nodesChan {
		fmt.Printf("收到节点: %v\n", node)
	}
	return err
}
