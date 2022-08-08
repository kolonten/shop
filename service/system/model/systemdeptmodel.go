package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemDeptModel = (*customSystemDeptModel)(nil)

type (
	// SystemDeptModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemDeptModel.
	SystemDeptModel interface {
		systemDeptModel
	}

	customSystemDeptModel struct {
		*defaultSystemDeptModel
	}
)

// NewSystemDeptModel returns a model for the database table.
func NewSystemDeptModel(conn sqlx.SqlConn, c cache.CacheConf) SystemDeptModel {
	return &customSystemDeptModel{
		defaultSystemDeptModel: newSystemDeptModel(conn, c),
	}
}
