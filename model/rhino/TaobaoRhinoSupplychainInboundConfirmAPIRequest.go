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
type TaobaoRhinoSupplychainInboundConfirmAPIRequest struct {
    model.Params
    // 入库单确认对象
    _clothingInboundConfirm   *ClothingInboundConfirmDto
}

// 初始化TaobaoRhinoSupplychainInboundConfirmAPIRequest对象
func NewTaobaoRhinoSupplychainInboundConfirmRequest() *TaobaoRhinoSupplychainInboundConfirmAPIRequest{
    return &TaobaoRhinoSupplychainInboundConfirmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRhinoSupplychainInboundConfirmAPIRequest) GetApiMethodName() string {
    return "taobao.rhino.supplychain.inbound.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRhinoSupplychainInboundConfirmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClothingInboundConfirm Setter
// 入库单确认对象
func (r *TaobaoRhinoSupplychainInboundConfirmAPIRequest) SetClothingInboundConfirm(_clothingInboundConfirm *ClothingInboundConfirmDto) error {
    r._clothingInboundConfirm = _clothingInboundConfirm
    r.Set("clothing_inbound_confirm", _clothingInboundConfirm)
    return nil
}

// ClothingInboundConfirm Getter
func (r TaobaoRhinoSupplychainInboundConfirmAPIRequest) GetClothingInboundConfirm() *ClothingInboundConfirmDto {
    return r._clothingInboundConfirm
}
