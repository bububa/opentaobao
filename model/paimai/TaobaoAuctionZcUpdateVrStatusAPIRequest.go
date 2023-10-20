package paimai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoauctionzcupdatevrstatusAPIRequest 如视VR更新活跃状态 API请求
// taobao.auction.zc.update.vr.status
//
// 如视VR更新活跃状态
type TaobaoauctionzcupdatevrstatusAPIRequest struct {
	model.Params
	// VR信息
	_message string
}

// NewTaobaoauctionzcupdatevrstatusRequest 初始化TaobaoauctionzcupdatevrstatusAPIRequest对象
func NewTaobaoauctionzcupdatevrstatusRequest() *TaobaoauctionzcupdatevrstatusAPIRequest {
	return &TaobaoauctionzcupdatevrstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoauctionzcupdatevrstatusAPIRequest) GetApiMethodName() string {
	return "taobao.auction.zc.update.vr.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoauctionzcupdatevrstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoauctionzcupdatevrstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMessage is Message Setter
// VR信息
func (r *TaobaoauctionzcupdatevrstatusAPIRequest) SetMessage(_message string) error {
	r._message = _message
	r.Set("message", _message)
	return nil
}

// GetMessage Message Getter
func (r TaobaoauctionzcupdatevrstatusAPIRequest) GetMessage() string {
	return r._message
}
