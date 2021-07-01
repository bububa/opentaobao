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
type TaobaoRhinoSupplychainClothingAdjustAPIRequest struct {
    model.Params
    // 库存盘点对象
    _param0   *MaterialInventoryAdjustDTO
}

// 初始化TaobaoRhinoSupplychainClothingAdjustAPIRequest对象
func NewTaobaoRhinoSupplychainClothingAdjustRequest() *TaobaoRhinoSupplychainClothingAdjustAPIRequest{
    return &TaobaoRhinoSupplychainClothingAdjustAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRhinoSupplychainClothingAdjustAPIRequest) GetApiMethodName() string {
    return "taobao.rhino.supplychain.clothing.adjust"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRhinoSupplychainClothingAdjustAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 库存盘点对象
func (r *TaobaoRhinoSupplychainClothingAdjustAPIRequest) SetParam0(_param0 *MaterialInventoryAdjustDTO) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoRhinoSupplychainClothingAdjustAPIRequest) GetParam0() *MaterialInventoryAdjustDTO {
    return r._param0
}
