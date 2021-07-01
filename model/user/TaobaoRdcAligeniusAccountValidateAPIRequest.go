package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdcAligeniusAccountValidateAPIRequest
AG商家账号校验 API请求
taobao.rdc.aligenius.account.validate

提供应对接AG的erp系统查询其旗下的商家是否为AG商家 */
type TaobaoRdcAligeniusAccountValidateAPIRequest struct {
	model.Params
}

// New
