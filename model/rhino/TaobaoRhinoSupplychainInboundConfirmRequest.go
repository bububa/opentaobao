package rhino

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
WMS003成衣入库确认 API请求
taobao.rhino.supplychain.inbound.confirm

【WMS003】【同步成衣入库完成信息】
*/
type TaobaoRhinoSupplychainInboundConfirmRequest struct {
    model.Params
    // 入库单确认对象
    clothingInboundConfirm   *ClothingInboundConfirmDto
}

// 初始化TaobaoRhinoSupplychainInboundConfirmRequest对象
func NewTaobaoRhinoSupplychainInboundConfirmRequest() *TaobaoRhinoSupplychainInboundConfirmRequest{
    return &TaobaoRhinoSupplychainInboundConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRhinoSupplychainInboundConfirmRequest) GetApiMethodName() string {
    return "taobao.rhino.supplychain.inbound.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRhinoSupplychainInboundConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClothingInboundConfirm Setter
// 入库单确认对象
func (r *TaobaoRhinoSupplychainInboundConfirmRequest) SetClothingInboundConfirm(clothingInboundConfirm *ClothingInboundConfirmDto) error {
    r.clothingInboundConfirm = clothingInboundConfirm
    r.Set("clothing_inbound_confirm", clothingInboundConfirm)
    return nil
}

// ClothingInboundConfirm Getter
func (r TaobaoRhinoSupplychainInboundConfirmRequest) GetClothingInboundConfirm() *ClothingInboundConfirmDto {
    return r.clothingInboundConfirm
}
