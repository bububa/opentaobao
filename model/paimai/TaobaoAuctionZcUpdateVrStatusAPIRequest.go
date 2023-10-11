package paimai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionZcUpdateVrStatusAPIRequest 如视VR更新活跃状态 API请求
// taobao.auction.zc.update.vr.status
//
// 如视VR更新活跃状态
type TaobaoAuctionZcUpdateVrStatusAPIRequest struct {
	model.Params
	// VR信息
	_message string
}

// NewTaobaoAuctionZcUpdateVrStatusRequest 初始化TaobaoAuctionZcUpdateVrStatusAPIRequest对象
func NewTaobaoAuctionZcUpdateVrStatusRequest() *TaobaoAuctionZcUpdateVrStatusAPIRequest {
	return &TaobaoAuctionZcUpdateVrStatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAuctionZcUpdateVrStatusAPIRequest) GetApiMethodName() string {
	return "taobao.auction.zc.update.vr.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAuctionZcUpdateVrStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAuctionZcUpdateVrStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMessage is Message Setter
// VR信息
func (r *TaobaoAuctionZcUpdateVrStatusAPIRequest) SetMessage(_message string) error {
	r._message = _message
	r.Set("message", _message)
	return nil
}

// GetMessage Message Getter
func (r TaobaoAuctionZcUpdateVrStatusAPIRequest) GetMessage() string {
	return r._message
}
