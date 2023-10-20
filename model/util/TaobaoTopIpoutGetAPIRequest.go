package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopipoutgetAPIRequest 获取开放平台出口IP段 API请求
// taobao.top.ipout.get
//
// 获取开放平台出口IP段
type TaobaotopipoutgetAPIRequest struct {
	model.Params
}

// NewTaobaotopipoutgetRequest 初始化TaobaotopipoutgetAPIRequest对象
func NewTaobaotopipoutgetRequest() *TaobaotopipoutgetAPIRequest {
	return &TaobaotopipoutgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotopipoutgetAPIRequest) GetApiMethodName() string {
	return "taobao.top.ipout.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotopipoutgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotopipoutgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
