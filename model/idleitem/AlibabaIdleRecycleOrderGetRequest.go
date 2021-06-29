package idleitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收订单查询V2 APIRequest
alibaba.idle.recycle.order.get

闲鱼回收业务中,外部回收商作为交易上买家,闲鱼用户下单后,需要回收商主动拉取交易订单
*/
type AlibabaIdleRecycleOrderGetRequest struct {
    model.Params

    // 订单号
    bizOrderId   int64 

}

func NewAlibabaIdleRecycleOrderGetRequest() *AlibabaIdleRecycleOrderGetRequest{
    return &AlibabaIdleRecycleOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRecycleOrderGetRequest) GetApiMethodName() string {
    return "alibaba.idle.recycle.order.get"
}

func (r AlibabaIdleRecycleOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRecycleOrderGetRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r AlibabaIdleRecycleOrderGetRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

