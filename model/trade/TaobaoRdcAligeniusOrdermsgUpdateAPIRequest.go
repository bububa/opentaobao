package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdcAligeniusOrdermsgUpdateAPIRequest
订单消息状态回传 API请求
taobao.rdc.aligenius.ordermsg.update

用于订单消息处理状态回传 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.ordermsg.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Oid Setter
// 子订单（消息中传的子订单）
func (r *TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) SetOid(_oid int64) error {
	r._oid = _oid
	r.Set("oid", _oid)
	return nil
}

// Get Oid Getter
func (r TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) GetOid() int64 {
	return r._oid
}

// Set is Status Setter
// 处理状态，1=成功，2=处理失败
func (r *TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) GetStatus() int64 {
	return r._status
}

// Set is Tid Setter
// 主订单（消息中传的主订单）
func (r *TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// Get Tid Getter
func (r TaobaoRdcAligeniusOrdermsgUpdateAPIRequest) GetTid() int64 {
	return r._tid
}
