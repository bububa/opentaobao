package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusOrdermsgUpdateAPIRequest 订单消息状态回传 API请求
// taobao.rdc.aligenius.ordermsg.update
//
// 用于订单消息处理状态回传
type TaobaoRdcAligeniusOrdermsgUpdateAPIRequest struct {
	model.Params
	// 子订单（消息中传的子订单）
	_oid int64
	// 处理状态，1=成功，2=处理失败
	_status int64
	// 主订单（消息中传的主订单）
	_tid int64
}

// NewTaobaoRdcAligeniusOrdermsgUpdateRequest 初始化TaobaoRdcAligeniusOrdermsgUpdateAPIRequest对象
func NewTaobaoRdcAligeniusOrdermsgUpdateRequest() *TaobaoRdcAligeniusOrdermsgUpdateAPIRequest {
	return &TaobaoRdcAligeniusOrdermsgUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) Reset() {
	r._oid = 0
	r._status = 0
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.ordermsg.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOid is Oid Setter
// 子订单（消息中传的子订单）
func (r *TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) SetOid(_oid int64) error {
	r._oid = _oid
	r.Set("oid", _oid)
	return nil
}

// GetOid Oid Getter
func (r TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) GetOid() int64 {
	return r._oid
}

// SetStatus is Status Setter
// 处理状态，1=成功，2=处理失败
func (r *TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) GetStatus() int64 {
	return r._status
}

// SetTid is Tid Setter
// 主订单（消息中传的主订单）
func (r *TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoRdcAligeniusOrdermsgUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRdcAligeniusOrdermsgUpdateRequest()
	},
}

// GetTaobaoRdcAligeniusOrdermsgUpdateRequest 从 sync.Pool 获取 TaobaoRdcAligeniusOrdermsgUpdateAPIRequest
func GetTaobaoRdcAligeniusOrdermsgUpdateAPIRequest() *TaobaoRdcAligeniusOrdermsgUpdateAPIRequest {
	return poolTaobaoRdcAligeniusOrdermsgUpdateAPIRequest.Get().(*TaobaoRdcAligeniusOrdermsgUpdateAPIRequest)
}

// ReleaseTaobaoRdcAligeniusOrdermsgUpdateAPIRequest 将 TaobaoRdcAligeniusOrdermsgUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoRdcAligeniusOrdermsgUpdateAPIRequest(v *TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) {
	v.Reset()
	poolTaobaoRdcAligeniusOrdermsgUpdateAPIRequest.Put(v)
}
