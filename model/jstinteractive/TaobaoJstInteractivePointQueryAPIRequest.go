package jstinteractive

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstInteractivePointQueryAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractivePointQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.point.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractivePointQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstInteractivePointQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoJstInteractivePointQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstInteractivePointQueryRequest()
	},
}

// GetTaobaoJstInteractivePointQueryRequest 从 sync.Pool 获取 TaobaoJstInteractivePointQueryAPIRequest
func GetTaobaoJstInteractivePointQueryAPIRequest() *TaobaoJstInteractivePointQueryAPIRequest {
	return poolTaobaoJstInteractivePointQueryAPIRequest.Get().(*TaobaoJstInteractivePointQueryAPIRequest)
}

// ReleaseTaobaoJstInteractivePointQueryAPIRequest 将 TaobaoJstInteractivePointQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstInteractivePointQueryAPIRequest(v *TaobaoJstInteractivePointQueryAPIRequest) {
	v.Reset()
	poolTaobaoJstInteractivePointQueryAPIRequest.Put(v)
}
