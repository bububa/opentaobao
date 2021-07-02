package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractivePointQueryAPIRequest 互动积分查询接口 API请求
// taobao.jst.interactive.point.query
//
// 查询用户的互动积分
type TaobaoJstInteractivePointQueryAPIRequest struct {
	model.Params
}

// NewTaobaoJstInteractivePointQueryRequest 初始化TaobaoJstInteractivePointQueryAPIRequest对象
func NewTaobaoJstInteractivePointQueryRequest() *TaobaoJstInteractivePointQueryAPIRequest {
	return &TaobaoJstInteractivePointQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractivePointQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.point.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractivePointQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
