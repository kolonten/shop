package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemDictModel = (*customSystemDictModel)(nil)

type (
	// SystemDictModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemDictModel.
	SystemDictModel interface {
		systemDictModel
	}

	customSystemDictModel struct {
		*defaultSystemDictModel
	}
)

// NewSystemDictModel returns a model for the database table.
func NewSystemDictModel(conn sqlx.SqlConn, c cache.CacheConf) SystemDictModel {
	return &customSystemDictModel{
		defaultSystemDictModel: newSystemDictModel(conn, c),
	}
}
