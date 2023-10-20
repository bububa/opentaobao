package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaohttpdnsgetAPIRequest TOPDNS配置 API请求
// taobao.httpdns.get
//
// 获取TOP DNS配置
type TaobaohttpdnsgetAPIRequest struct {
	model.Params
}

// NewTaobaohttpdnsgetRequest 初始化TaobaohttpdnsgetAPIRequest对象
func NewTaobaohttpdnsgetRequest() *TaobaohttpdnsgetAPIRequest {
	return &TaobaohttpdnsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaohttpdnsgetAPIRequest) GetApiMethodName() string {
	return "taobao.httpdns.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaohttpdnsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaohttpdnsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
