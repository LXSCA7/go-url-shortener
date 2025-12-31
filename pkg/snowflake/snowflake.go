package snowflake

import (
	"fmt"
	"sync"
	"time"
)

type Node struct {
	mu        sync.Mutex
	timestamp int64
	nodeID    int64
	step      int64
}

const (
	nodeBits  = 10
	stepBits  = 12
	nodeMax   = -1 ^ (-1 << nodeBits)
	stepMax   = -1 ^ (-1 << stepBits)
	timeShift = nodeBits + stepBits
	nodeShift = stepBits
	epoch     = 1735689600000 // 01.01.2025
)

func NewNode(nodeID int64) (*Node, error) {
	if nodeID < 0 || nodeID > nodeMax {
		return &Node{}, fmt.Errorf("invalid node ID. the node ID needs to be between 1 and 1023")
	}

	return &Node{
		timestamp: 0,
		nodeID:    nodeID,
		step:      0,
	}, nil
}

func (n *Node) Generate() int64 {
	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Now().UnixMilli()
	if now == n.timestamp {
		n.step = (n.step + 1) & stepMax

		if n.step == 0 {
			for now <= n.timestamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		n.step = 0
	}

	n.timestamp = now

	return ((now - epoch) << timeShift) | (n.nodeID << nodeShift) | n.step
}
