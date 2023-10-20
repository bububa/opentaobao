package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaordcaligeniusordermsgupdateAPIRequest 订单消息状态回传 API请求
// taobao.rdc.aligenius.ordermsg.update
//
// 用于订单消息处理状态回传
type TaobaordcaligeniusordermsgupdateAPIRequest struct {
	model.Params
	// 子订单（消息中传的子订单）
	_oid int64
	// 处理状态，1=成功，2=处理失败
	_status int64
	// 主订单（消息中传的主订单）
	_tid int64
}

// NewTaobaordcaligeniusordermsgupdateRequest 初始化TaobaordcaligeniusordermsgupdateAPIRequest对象
func NewTaobaordcaligeniusordermsgupdateRequest() *TaobaordcaligeniusordermsgupdateAPIRequest {
	return &TaobaordcaligeniusordermsgupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaordcaligeniusordermsgupdateAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.ordermsg.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaordcaligeniusordermsgupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaordcaligeniusordermsgupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOid is Oid Setter
// 子订单（消息中传的子订单）
func (r *TaobaordcaligeniusordermsgupdateAPIRequest) SetOid(_oid int64) error {
	r._oid = _oid
	r.Set("oid", _oid)
	return nil
}

// GetOid Oid Getter
func (r TaobaordcaligeniusordermsgupdateAPIRequest) GetOid() int64 {
	return r._oid
}

// SetStatus is Status Setter
// 处理状态，1=成功，2=处理失败
func (r *TaobaordcaligeniusordermsgupdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaordcaligeniusordermsgupdateAPIRequest) GetStatus() int64 {
	return r._status
}

// SetTid is Tid Setter
// 主订单（消息中传的主订单）
func (r *TaobaordcaligeniusordermsgupdateAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaordcaligeniusordermsgupdateAPIRequest) GetTid() int64 {
	return r._tid
}
