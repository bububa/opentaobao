package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopIpoutGetAPIRequest 获取开放平台出口IP段 API请求
// taobao.top.ipout.get
//
// 获取开放平台出口IP段
type TaobaoTopIpoutGetAPIRequest struct {
	model.Params
}

// NewTaobaoTopIpoutGetRequest 初始化TaobaoTopIpoutGetAPIRequest对象
func NewTaobaoTopIpoutGetRequest() *TaobaoTopIpoutGetAPIRequest {
	return &TaobaoTopIpoutGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopIpoutGetAPIRequest) GetApiMethodName() string {
	return "taobao.top.ipout.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopIpoutGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopIpoutGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
