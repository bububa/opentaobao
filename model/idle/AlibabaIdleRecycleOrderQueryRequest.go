package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收订单查询V1 APIRequest
alibaba.idle.recycle.order.query

查询回收订单
*/
type AlibabaIdleRecycleOrderQueryRequest struct {
    model.Params

    // 订单号
    bizOrderId   int64 

}

func NewAlibabaIdleRecycleOrderQueryRequest() *AlibabaIdleRecycleOrderQueryRequest{
    return &AlibabaIdleRecycleOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRecycleOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.recycle.order.query"
}

func (r AlibabaIdleRecycleOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRecycleOrderQueryRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r AlibabaIdleRecycleOrderQueryRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

