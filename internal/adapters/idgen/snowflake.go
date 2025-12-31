package idgen

import (
	"github.com/LXSCA7/go-url-shortener/internal/core/ports"
	"github.com/LXSCA7/go-url-shortener/pkg/snowflake"
)

type SnowflakeAdapter struct {
	node *snowflake.Node
}

func NewSnowflakeIDGen(nodeID int64) (*SnowflakeAdapter, error) {
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		return nil, err
	}

	return &SnowflakeAdapter{
		node: node,
	}, nil
}

var _ ports.IDGenerator = (*SnowflakeAdapter)(nil)

func (s *SnowflakeAdapter) Generate() int64 {
	return s.node.Generate()
}
