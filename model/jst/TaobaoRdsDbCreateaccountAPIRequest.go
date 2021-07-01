package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdsDbCreateaccountAPIRequest
rds创建数据库账户 API请求
taobao.rds.db.createaccount

rds创建数据库账户 */
type TaobaoRdsDbCreateaccountAPIRequest struct {
	model.Params
	// 入参对象
	_param0 *RequestDbAccountModel
}

// New
