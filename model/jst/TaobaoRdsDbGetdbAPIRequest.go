package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdsDbGetdbAPIRequest
rds获取RDS的DB API请求
taobao.rds.db.getdb

rds获取RDS的DB */
type TaobaoRdsDbGetdbAPIRequest struct {
	model.Params
	// 账户名
	_accountName string
	// 实例名
	_instanceName string
}

// New
