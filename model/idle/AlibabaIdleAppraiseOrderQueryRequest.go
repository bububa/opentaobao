package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼验货担保订单详情查询V1 APIRequest
alibaba.idle.appraise.order.query

鉴定商调用该接口获取订单状态
*/
type AlibabaIdleAppraiseOrderQueryRequest struct {
    model.Params

    // orderId
    bizOrderId   int64 

}

func NewAlibabaIdleAppraiseOrderQueryRequest() *AlibabaIdleAppraiseOrderQueryRequest{
    return &AlibabaIdleAppraiseOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleAppraiseOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.appraise.order.query"
}

func (r AlibabaIdleAppraiseOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleAppraiseOrderQueryRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r AlibabaIdleAppraiseOrderQueryRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

