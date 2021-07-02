package rhino

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRhinoSupplychainClothingAdjustAPIRequest 同步成衣仓盘点数据 API请求
// taobao.rhino.supplychain.clothing.adjust
//
// 同步成衣仓盘点数据
type TaobaoRhinoSupplychainClothingAdjustAPIRequest struct {
	model.Params
	// 库存盘点对象
	_param0 *MaterialInventoryAdjustDto
}

// NewTaobaoRhinoSupplychainClothingAdjustRequest 初始化TaobaoRhinoSupplychainClothingAdjustAPIRequest对象
func NewTaobaoRhinoSupplychainClothingAdjustRequest() *TaobaoRhinoSupplychainClothingAdjustAPIRequest {
	return &TaobaoRhinoSupplychainClothingAdjustAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRhinoSupplychainClothingAdjustAPIRequest) GetApiMethodName() string {
	return "taobao.rhino.supplychain.clothing.adjust"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRhinoSupplychainClothingAdjustAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 库存盘点对象
func (r *TaobaoRhinoSupplychainClothingAdjustAPIRequest) SetParam0(_param0 *MaterialInventoryAdjustDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TaobaoRhinoSupplychainClothingAdjustAPIRequest) GetParam0() *MaterialInventoryAdjustDto {
	return r._param0
}
