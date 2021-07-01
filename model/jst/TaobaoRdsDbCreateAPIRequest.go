package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdsDbCreateAPIRequest
rds创建数据库 API请求
taobao.rds.db.create

在rds实例里创建数据库 */
type TaobaoRdsDbCreateAPIRequest struct {
	model.Params
	// rds的实例名
	_instanceName string
	// 数据库名
	_dbName string
	// 已存在账号名
	_accountName string
}

// New
