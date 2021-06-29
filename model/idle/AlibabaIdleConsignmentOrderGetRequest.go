package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼帮卖订单查询 APIRequest
alibaba.idle.consignment.order.get

闲鱼帮卖服务商以闲鱼交易买家身份查询订单信息
*/
type AlibabaIdleConsignmentOrderGetRequest struct {
    model.Params

    // 闲鱼订单ID
    bizOrderId   int64 

}

func NewAlibabaIdleConsignmentOrderGetRequest() *AlibabaIdleConsignmentOrderGetRequest{
    return &AlibabaIdleConsignmentOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleConsignmentOrderGetRequest) GetApiMethodName() string {
    return "alibaba.idle.consignment.order.get"
}

func (r AlibabaIdleConsignmentOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleConsignmentOrderGetRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r AlibabaIdleConsignmentOrderGetRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

