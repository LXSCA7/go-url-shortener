package idgen

import (
	"github.com/LXSCA7/go-url-shortener/internal/core/ports"
	"github.com/LXSCA7/go-url-shortener/pkg/snowflake"
)

type SnowflakeAdapter struct {
	node *snowflake.Node
}

var _ ports.IDGenerator = (*SnowflakeAdapter)(nil)

func (s *SnowflakeAdapter) Generate() int64 {
	return s.node.Generate()
}
