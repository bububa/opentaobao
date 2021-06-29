package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼寄卖V2订单查询 APIRequest
alibaba.idle.consignmentii.order.get

闲鱼寄卖V2服务商以闲鱼交易买家身份查询订单信息
*/
type AlibabaIdleConsignmentiiOrderGetRequest struct {
    model.Params

    // 闲鱼订单ID
    bizOrderId   int64 

}

func NewAlibabaIdleConsignmentiiOrderGetRequest() *AlibabaIdleConsignmentiiOrderGetRequest{
    return &AlibabaIdleConsignmentiiOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleConsignmentiiOrderGetRequest) GetApiMethodName() string {
    return "alibaba.idle.consignmentii.order.get"
}

func (r AlibabaIdleConsignmentiiOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleConsignmentiiOrderGetRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r AlibabaIdleConsignmentiiOrderGetRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

