package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
盒马帮采购确认订单接口 APIRequest
alibaba.wdk.wholesale.order.commit

盒马帮采购确认订单接口
*/
type AlibabaWdkWholesaleOrderCommitRequest struct {
    model.Params

    // 采购单信息
    orderCommitReq   *OrderCommitReq 

}

func NewAlibabaWdkWholesaleOrderCommitRequest() *AlibabaWdkWholesaleOrderCommitRequest{
    return &AlibabaWdkWholesaleOrderCommitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkWholesaleOrderCommitRequest) GetApiMethodName() string {
    return "alibaba.wdk.wholesale.order.commit"
}

func (r AlibabaWdkWholesaleOrderCommitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkWholesaleOrderCommitRequest) SetOrderCommitReq(orderCommitReq *OrderCommitReq) error {
    r.orderCommitReq = orderCommitReq
    r.Set("order_commit_req", orderCommitReq)
    return nil
}

func (r AlibabaWdkWholesaleOrderCommitRequest) GetOrderCommitReq() *OrderCommitReq {
    return r.orderCommitReq
}

