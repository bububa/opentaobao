package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstinteractivepointqueryAPIRequest 互动积分查询接口 API请求
// taobao.jst.interactive.point.query
//
// 查询用户的互动积分
type TaobaojstinteractivepointqueryAPIRequest struct {
	model.Params
}

// NewTaobaojstinteractivepointqueryRequest 初始化TaobaojstinteractivepointqueryAPIRequest对象
func NewTaobaojstinteractivepointqueryRequest() *TaobaojstinteractivepointqueryAPIRequest {
	return &TaobaojstinteractivepointqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstinteractivepointqueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.point.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstinteractivepointqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstinteractivepointqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
