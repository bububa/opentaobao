package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCntecItemChangeMessageAPIRequest 商品变更消息 API请求
// cainiao.cntec.item.change.message
//
// 供货商商品信息变更消息
type CainiaoCntecItemChangeMessageAPIRequest struct {
	model.Params
	// 供应商商品变更信息
	_itemChangeMessage *SupplyItemChangeMessage
}

// NewCainiaoCntecItemChangeMessageRequest 初始化CainiaoCntecItemChangeMessageAPIRequest对象
func NewCainiaoCntecItemChangeMessageRequest() *CainiaoCntecItemChangeMessageAPIRequest {
	return &CainiaoCntecItemChangeMessageAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoCntecItemChangeMessageAPIRequest) Reset() {
	r._itemChangeMessage = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCntecItemChangeMessageAPIRequest) GetApiMethodName() string {
	return "cainiao.cntec.item.change.message"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCntecItemChangeMessageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCntecItemChangeMessageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemChangeMessage is ItemChangeMessage Setter
// 供应商商品变更信息
func (r *CainiaoCntecItemChangeMessageAPIRequest) SetItemChangeMessage(_itemChangeMessage *SupplyItemChangeMessage) error {
	r._itemChangeMessage = _itemChangeMessage
	r.Set("item_change_message", _itemChangeMessage)
	return nil
}

// GetItemChangeMessage ItemChangeMessage Getter
func (r CainiaoCntecItemChangeMessageAPIRequest) GetItemChangeMessage() *SupplyItemChangeMessage {
	return r._itemChangeMessage
}

var poolCainiaoCntecItemChangeMessageAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoCntecItemChangeMessageRequest()
	},
}

// GetCainiaoCntecItemChangeMessageRequest 从 sync.Pool 获取 CainiaoCntecItemChangeMessageAPIRequest
func GetCainiaoCntecItemChangeMessageAPIRequest() *CainiaoCntecItemChangeMessageAPIRequest {
	return poolCainiaoCntecItemChangeMessageAPIRequest.Get().(*CainiaoCntecItemChangeMessageAPIRequest)
}

// ReleaseCainiaoCntecItemChangeMessageAPIRequest 将 CainiaoCntecItemChangeMessageAPIRequest 放入 sync.Pool
func ReleaseCainiaoCntecItemChangeMessageAPIRequest(v *CainiaoCntecItemChangeMessageAPIRequest) {
	v.Reset()
	poolCainiaoCntecItemChangeMessageAPIRequest.Put(v)
}
