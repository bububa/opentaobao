package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotimegetAPIRequest 获取淘宝系统当前时间 API请求
// taobao.time.get
//
// 获取淘宝系统当前时间
type TaobaotimegetAPIRequest struct {
	model.Params
}

// NewTaobaotimegetRequest 初始化TaobaotimegetAPIRequest对象
func NewTaobaotimegetRequest() *TaobaotimegetAPIRequest {
	return &TaobaotimegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotimegetAPIRequest) GetApiMethodName() string {
	return "taobao.time.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotimegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotimegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
