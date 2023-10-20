package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusAccountValidateAPIRequest AG商家账号校验 API请求
// taobao.rdc.aligenius.account.validate
//
// 提供应对接AG的erp系统查询其旗下的商家是否为AG商家
type TaobaoRdcAligeniusAccountValidateAPIRequest struct {
	model.Params
}

// NewTaobaoRdcAligeniusAccountValidateRequest 初始化TaobaoRdcAligeniusAccountValidateAPIRequest对象
func NewTaobaoRdcAligeniusAccountValidateRequest() *TaobaoRdcAligeniusAccountValidateAPIRequest {
	return &TaobaoRdcAligeniusAccountValidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusAccountValidateAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.account.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusAccountValidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdcAligeniusAccountValidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}
