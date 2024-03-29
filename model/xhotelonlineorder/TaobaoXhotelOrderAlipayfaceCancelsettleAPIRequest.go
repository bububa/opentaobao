package xhotelonlineorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest 信用住订单取消结算接口 API请求
// taobao.xhotel.order.alipayface.cancelsettle
//
// 信用住订单由于客人为出现等原因，最终取消结算,一定要在结算后2个小时之内调用，否则不会成功。
type TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest struct {
	model.Params
	// 取消结账的原因
	_reason string
	// 外部订单号，和tid二选一必填
	_outId string
	// 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
	_tid int64
}

// NewTaobaoXhotelOrderAlipayfaceCancelsettleRequest 初始化TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest对象
func NewTaobaoXhotelOrderAlipayfaceCancelsettleRequest() *TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest {
	return &TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest) Reset() {
	r._reason = ""
	r._outId = ""
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.alipayface.cancelsettle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReason is Reason Setter
// 取消结账的原因
func (r *TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest) GetReason() string {
	return r._reason
}

// SetOutId is OutId Setter
// 外部订单号，和tid二选一必填
func (r *TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("out_id", _outId)
	return nil
}

// GetOutId OutId Getter
func (r TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest) GetOutId() string {
	return r._outId
}

// SetTid is Tid Setter
// 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
func (r *TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelOrderAlipayfaceCancelsettleRequest()
	},
}

// GetTaobaoXhotelOrderAlipayfaceCancelsettleRequest 从 sync.Pool 获取 TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest
func GetTaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest() *TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest {
	return poolTaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest.Get().(*TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest)
}

// ReleaseTaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest 将 TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest(v *TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest) {
	v.Reset()
	poolTaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest.Put(v)
}
