package rhino

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【WMS005】接收成衣捡配完成通知 APIRequest
taobao.rhino.supplychain.outbound.pickingcomplete

接收成衣捡配完成通知,WMS005
*/
type TaobaoRhinoSupplychainOutboundPickingcompleteRequest struct {
    model.Params

    // 捡配完成消息
    param0   *PickingCompleteMsg 

}

func NewTaobaoRhinoSupplychainOutboundPickingcompleteRequest() *TaobaoRhinoSupplychainOutboundPickingcompleteRequest{
    return &TaobaoRhinoSupplychainOutboundPickingcompleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRhinoSupplychainOutboundPickingcompleteRequest) GetApiMethodName() string {
    return "taobao.rhino.supplychain.outbound.pickingcomplete"
}

func (r TaobaoRhinoSupplychainOutboundPickingcompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRhinoSupplychainOutboundPickingcompleteRequest) SetParam0(param0 *PickingCompleteMsg) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoRhinoSupplychainOutboundPickingcompleteRequest) GetParam0() *PickingCompleteMsg {
    return r.param0
}

