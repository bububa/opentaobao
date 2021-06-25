package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品变更消息 APIRequest
cainiao.cntec.item.change.message

供货商商品信息变更消息
*/
type CainiaoCntecItemChangeMessageRequest struct {
    model.Params

    // 供应商商品变更信息
    itemChangeMessage   *SupplyItemChangeMessage 

}

func NewCainiaoCntecItemChangeMessageRequest() *CainiaoCntecItemChangeMessageRequest{
    return &CainiaoCntecItemChangeMessageRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCntecItemChangeMessageRequest) GetApiMethodName() string {
    return "cainiao.cntec.item.change.message"
}

func (r CainiaoCntecItemChangeMessageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCntecItemChangeMessageRequest) SetItemChangeMessage(itemChangeMessage *SupplyItemChangeMessage) error {
    r.itemChangeMessage = itemChangeMessage
    r.Set("item_change_message", itemChangeMessage)
    return nil
}

func (r CainiaoCntecItemChangeMessageRequest) GetItemChangeMessage() *SupplyItemChangeMessage {
    return r.itemChangeMessage
}

