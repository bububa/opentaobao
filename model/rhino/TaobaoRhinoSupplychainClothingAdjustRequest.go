package rhino

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步成衣仓盘点数据 APIRequest
taobao.rhino.supplychain.clothing.adjust

同步成衣仓盘点数据
*/
type TaobaoRhinoSupplychainClothingAdjustRequest struct {
    model.Params

    // 库存盘点对象
    param0   *MaterialInventoryAdjustDto 

}

func NewTaobaoRhinoSupplychainClothingAdjustRequest() *TaobaoRhinoSupplychainClothingAdjustRequest{
    return &TaobaoRhinoSupplychainClothingAdjustRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRhinoSupplychainClothingAdjustRequest) GetApiMethodName() string {
    return "taobao.rhino.supplychain.clothing.adjust"
}

func (r TaobaoRhinoSupplychainClothingAdjustRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRhinoSupplychainClothingAdjustRequest) SetParam0(param0 *MaterialInventoryAdjustDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoRhinoSupplychainClothingAdjustRequest) GetParam0() *MaterialInventoryAdjustDto {
    return r.param0
}

