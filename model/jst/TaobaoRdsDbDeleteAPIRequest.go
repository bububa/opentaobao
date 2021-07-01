package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdsDbDeleteAPIRequest
RDS数据库删除 API请求
taobao.rds.db.delete

通过api删除用户RDS的数据库 */
type TaobaoRdsDbDeleteAPIRequest struct {
	model.Params
	// rds的实例名
	_instanceName string
	// 数据库的name，可以通过 taobao.rds.db.get 获取
	_dbName string
}

// New
