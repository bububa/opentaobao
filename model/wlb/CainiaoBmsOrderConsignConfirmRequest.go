package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
BMS出库通知 APIRequest
cainiao.bms.order.consign.confirm

BMS出库后，通知ISV
*/
type CainiaoBmsOrderConsignConfirmRequest struct {
    model.Params

    // 通知消息主体
    content   *BmsConsignOrderConfirm 

}

func NewCainiaoBmsOrderConsignConfirmRequest() *CainiaoBmsOrderConsignConfirmRequest{
    return &CainiaoBmsOrderConsignConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoBmsOrderConsignConfirmRequest) GetApiMethodName() string {
    return "cainiao.bms.order.consign.confirm"
}

func (r CainiaoBmsOrderConsignConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoBmsOrderConsignConfirmRequest) SetContent(content *BmsConsignOrderConfirm) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r CainiaoBmsOrderConsignConfirmRequest) GetContent() *BmsConsignOrderConfirm {
    return r.content
}

