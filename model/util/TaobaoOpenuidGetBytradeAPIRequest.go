package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenuidGetBytradeAPIRequest 通过订单获取对应买家的openUID API请求
// taobao.openuid.get.bytrade
//
// 通过订单获取对应买家的openUID,需要卖家授权
type TaobaoOpenuidGetBytradeAPIRequest struct {
	model.Params
	// 订单ID
	_tid int64
}

// NewTaobaoOpenuidGetBytradeRequest 初始化TaobaoOpenuidGetBytradeAPIRequest对象
func NewTaobaoOpenuidGetBytradeRequest() *TaobaoOpenuidGetBytradeAPIRequest {
	return &TaobaoOpenuidGetBytradeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenuidGetBytradeAPIRequest) Reset() {
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenuidGetBytradeAPIRequest) GetApiMethodName() string {
	return "taobao.openuid.get.bytrade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenuidGetBytradeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenuidGetBytradeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 订单ID
func (r *TaobaoOpenuidGetBytradeAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoOpenuidGetBytradeAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoOpenuidGetBytradeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenuidGetBytradeRequest()
	},
}

// GetTaobaoOpenuidGetBytradeRequest 从 sync.Pool 获取 TaobaoOpenuidGetBytradeAPIRequest
func GetTaobaoOpenuidGetBytradeAPIRequest() *TaobaoOpenuidGetBytradeAPIRequest {
	return poolTaobaoOpenuidGetBytradeAPIRequest.Get().(*TaobaoOpenuidGetBytradeAPIRequest)
}

// ReleaseTaobaoOpenuidGetBytradeAPIRequest 将 TaobaoOpenuidGetBytradeAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenuidGetBytradeAPIRequest(v *TaobaoOpenuidGetBytradeAPIRequest) {
	v.Reset()
	poolTaobaoOpenuidGetBytradeAPIRequest.Put(v)
}
