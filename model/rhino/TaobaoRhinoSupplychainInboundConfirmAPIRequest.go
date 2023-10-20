package rhino

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRhinoSupplychainInboundConfirmAPIRequest WMS003成衣入库确认 API请求
// taobao.rhino.supplychain.inbound.confirm
//
// 【WMS003】【同步成衣入库完成信息】
type TaobaoRhinoSupplychainInboundConfirmAPIRequest struct {
	model.Params
	// 入库单确认对象
	_clothingInboundConfirm *ClothingInboundConfirmDto
}

// NewTaobaoRhinoSupplychainInboundConfirmRequest 初始化TaobaoRhinoSupplychainInboundConfirmAPIRequest对象
func NewTaobaoRhinoSupplychainInboundConfirmRequest() *TaobaoRhinoSupplychainInboundConfirmAPIRequest {
	return &TaobaoRhinoSupplychainInboundConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRhinoSupplychainInboundConfirmAPIRequest) Reset() {
	r._clothingInboundConfirm = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRhinoSupplychainInboundConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.rhino.supplychain.inbound.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRhinoSupplychainInboundConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRhinoSupplychainInboundConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClothingInboundConfirm is ClothingInboundConfirm Setter
// 入库单确认对象
func (r *TaobaoRhinoSupplychainInboundConfirmAPIRequest) SetClothingInboundConfirm(_clothingInboundConfirm *ClothingInboundConfirmDto) error {
	r._clothingInboundConfirm = _clothingInboundConfirm
	r.Set("clothing_inbound_confirm", _clothingInboundConfirm)
	return nil
}

// GetClothingInboundConfirm ClothingInboundConfirm Getter
func (r TaobaoRhinoSupplychainInboundConfirmAPIRequest) GetClothingInboundConfirm() *ClothingInboundConfirmDto {
	return r._clothingInboundConfirm
}

var poolTaobaoRhinoSupplychainInboundConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRhinoSupplychainInboundConfirmRequest()
	},
}

// GetTaobaoRhinoSupplychainInboundConfirmRequest 从 sync.Pool 获取 TaobaoRhinoSupplychainInboundConfirmAPIRequest
func GetTaobaoRhinoSupplychainInboundConfirmAPIRequest() *TaobaoRhinoSupplychainInboundConfirmAPIRequest {
	return poolTaobaoRhinoSupplychainInboundConfirmAPIRequest.Get().(*TaobaoRhinoSupplychainInboundConfirmAPIRequest)
}

// ReleaseTaobaoRhinoSupplychainInboundConfirmAPIRequest 将 TaobaoRhinoSupplychainInboundConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoRhinoSupplychainInboundConfirmAPIRequest(v *TaobaoRhinoSupplychainInboundConfirmAPIRequest) {
	v.Reset()
	poolTaobaoRhinoSupplychainInboundConfirmAPIRequest.Put(v)
}
