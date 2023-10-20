package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocntecitemchangemessageAPIRequest 商品变更消息 API请求
// cainiao.cntec.item.change.message
//
// 供货商商品信息变更消息
type CainiaocntecitemchangemessageAPIRequest struct {
	model.Params
	// 供应商商品变更信息
	_itemChangeMessage *SupplyItemChangeMessage
}

// NewCainiaocntecitemchangemessageRequest 初始化CainiaocntecitemchangemessageAPIRequest对象
func NewCainiaocntecitemchangemessageRequest() *CainiaocntecitemchangemessageAPIRequest {
	return &CainiaocntecitemchangemessageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocntecitemchangemessageAPIRequest) GetApiMethodName() string {
	return "cainiao.cntec.item.change.message"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocntecitemchangemessageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocntecitemchangemessageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemChangeMessage is ItemChangeMessage Setter
// 供应商商品变更信息
func (r *CainiaocntecitemchangemessageAPIRequest) SetItemChangeMessage(_itemChangeMessage *SupplyItemChangeMessage) error {
	r._itemChangeMessage = _itemChangeMessage
	r.Set("item_change_message", _itemChangeMessage)
	return nil
}

// GetItemChangeMessage ItemChangeMessage Getter
func (r CainiaocntecitemchangemessageAPIRequest) GetItemChangeMessage() *SupplyItemChangeMessage {
	return r._itemChangeMessage
}
