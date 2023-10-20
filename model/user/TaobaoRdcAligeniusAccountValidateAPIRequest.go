package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaordcaligeniusaccountvalidateAPIRequest AG商家账号校验 API请求
// taobao.rdc.aligenius.account.validate
//
// 提供应对接AG的erp系统查询其旗下的商家是否为AG商家
type TaobaordcaligeniusaccountvalidateAPIRequest struct {
	model.Params
}

// NewTaobaordcaligeniusaccountvalidateRequest 初始化TaobaordcaligeniusaccountvalidateAPIRequest对象
func NewTaobaordcaligeniusaccountvalidateRequest() *TaobaordcaligeniusaccountvalidateAPIRequest {
	return &TaobaordcaligeniusaccountvalidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaordcaligeniusaccountvalidateAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.account.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaordcaligeniusaccountvalidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaordcaligeniusaccountvalidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}
