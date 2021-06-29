package rhino

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
WMS003成衣入库确认 APIRequest
taobao.rhino.supplychain.inbound.confirm

【WMS003】【同步成衣入库完成信息】
*/
type TaobaoRhinoSupplychainInboundConfirmRequest struct {
    model.Params

    // 入库单确认对象
    clothingInboundConfirm   *ClothingInboundConfirmDto 

}

func NewTaobaoRhinoSupplychainInboundConfirmRequest() *TaobaoRhinoSupplychainInboundConfirmRequest{
    return &TaobaoRhinoSupplychainInboundConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRhinoSupplychainInboundConfirmRequest) GetApiMethodName() string {
    return "taobao.rhino.supplychain.inbound.confirm"
}

func (r TaobaoRhinoSupplychainInboundConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRhinoSupplychainInboundConfirmRequest) SetClothingInboundConfirm(clothingInboundConfirm *ClothingInboundConfirmDto) error {
    r.clothingInboundConfirm = clothingInboundConfirm
    r.Set("clothing_inbound_confirm", clothingInboundConfirm)
    return nil
}

func (r TaobaoRhinoSupplychainInboundConfirmRequest) GetClothingInboundConfirm() *ClothingInboundConfirmDto {
    return r.clothingInboundConfirm
}

