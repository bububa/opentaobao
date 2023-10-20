package rhino

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorhinosupplychaininboundconfirmAPIRequest WMS003成衣入库确认 API请求
// taobao.rhino.supplychain.inbound.confirm
//
// 【WMS003】【同步成衣入库完成信息】
type TaobaorhinosupplychaininboundconfirmAPIRequest struct {
	model.Params
	// 入库单确认对象
	_clothingInboundConfirm *ClothingInboundConfirmDto
}

// NewTaobaorhinosupplychaininboundconfirmRequest 初始化TaobaorhinosupplychaininboundconfirmAPIRequest对象
func NewTaobaorhinosupplychaininboundconfirmRequest() *TaobaorhinosupplychaininboundconfirmAPIRequest {
	return &TaobaorhinosupplychaininboundconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorhinosupplychaininboundconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.rhino.supplychain.inbound.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorhinosupplychaininboundconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorhinosupplychaininboundconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClothingInboundConfirm is ClothingInboundConfirm Setter
// 入库单确认对象
func (r *TaobaorhinosupplychaininboundconfirmAPIRequest) SetClothingInboundConfirm(_clothingInboundConfirm *ClothingInboundConfirmDto) error {
	r._clothingInboundConfirm = _clothingInboundConfirm
	r.Set("clothing_inbound_confirm", _clothingInboundConfirm)
	return nil
}

// GetClothingInboundConfirm ClothingInboundConfirm Getter
func (r TaobaorhinosupplychaininboundconfirmAPIRequest) GetClothingInboundConfirm() *ClothingInboundConfirmDto {
	return r._clothingInboundConfirm
}
