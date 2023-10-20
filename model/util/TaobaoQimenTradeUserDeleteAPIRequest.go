package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenTradeUserDeleteAPIRequest 删除奇门订单链路用户 API请求
// taobao.qimen.trade.user.delete
//
// 删除奇门订单链路用户
type TaobaoQimenTradeUserDeleteAPIRequest struct {
	model.Params
}

// NewTaobaoQimenTradeUserDeleteRequest 初始化TaobaoQimenTradeUserDeleteAPIRequest对象
func NewTaobaoQimenTradeUserDeleteRequest() *TaobaoQimenTradeUserDeleteAPIRequest {
	return &TaobaoQimenTradeUserDeleteAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenTradeUserDeleteAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenTradeUserDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.trade.user.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenTradeUserDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenTradeUserDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoQimenTradeUserDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenTradeUserDeleteRequest()
	},
}

// GetTaobaoQimenTradeUserDeleteRequest 从 sync.Pool 获取 TaobaoQimenTradeUserDeleteAPIRequest
func GetTaobaoQimenTradeUserDeleteAPIRequest() *TaobaoQimenTradeUserDeleteAPIRequest {
	return poolTaobaoQimenTradeUserDeleteAPIRequest.Get().(*TaobaoQimenTradeUserDeleteAPIRequest)
}

// ReleaseTaobaoQimenTradeUserDeleteAPIRequest 将 TaobaoQimenTradeUserDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenTradeUserDeleteAPIRequest(v *TaobaoQimenTradeUserDeleteAPIRequest) {
	v.Reset()
	poolTaobaoQimenTradeUserDeleteAPIRequest.Put(v)
}
