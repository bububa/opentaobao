package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojdshlusergetAPIRequest 订单全链路用户信息获取 API请求
// taobao.jds.hluser.get
//
// 订单全链路用户信息获取
type TaobaojdshlusergetAPIRequest struct {
	model.Params
}

// NewTaobaojdshlusergetRequest 初始化TaobaojdshlusergetAPIRequest对象
func NewTaobaojdshlusergetRequest() *TaobaojdshlusergetAPIRequest {
	return &TaobaojdshlusergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojdshlusergetAPIRequest) GetApiMethodName() string {
	return "taobao.jds.hluser.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojdshlusergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojdshlusergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
