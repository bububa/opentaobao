package rhino

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步成衣仓盘点数据 API请求
taobao.rhino.supplychain.clothing.adjust

同步成衣仓盘点数据
*/
type TaobaoRhinoSupplychainClothingAdjustRequest struct {
    model.Params
    // 库存盘点对象
    param0   *MaterialInventoryAdjustDto
}

// 初始化TaobaoRhinoSupplychainClothingAdjustRequest对象
func NewTaobaoRhinoSupplychainClothingAdjustRequest() *TaobaoRhinoSupplychainClothingAdjustRequest{
    return &TaobaoRhinoSupplychainClothingAdjustRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRhinoSupplychainClothingAdjustRequest) GetApiMethodName() string {
    return "taobao.rhino.supplychain.clothing.adjust"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRhinoSupplychainClothingAdjustRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 库存盘点对象
func (r *TaobaoRhinoSupplychainClothingAdjustRequest) SetParam0(param0 *MaterialInventoryAdjustDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TaobaoRhinoSupplychainClothingAdjustRequest) GetParam0() *MaterialInventoryAdjustDto {
    return r.param0
}
