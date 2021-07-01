package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdsDbGetAPIRequest
查询rds下的数据库 API请求
taobao.rds.db.get

查询rds实例下的数据库 */
type TaobaoRdsDbGetAPIRequest struct {
	model.Params
	// rds的实例名
	_instanceName string
	// 数据库状态，默认值1
	_dbStatus int64
}

// New
