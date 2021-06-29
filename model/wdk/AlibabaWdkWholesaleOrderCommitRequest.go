package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
盒马帮采购确认订单接口 API请求
alibaba.wdk.wholesale.order.commit

盒马帮采购确认订单接口
*/
type AlibabaWdkWholesaleOrderCommitRequest struct {
    model.Params
    // 采购单信息
    orderCommitReq   *OrderCommitReq
}

// 初始化AlibabaWdkWholesaleOrderCommitRequest对象
func NewAlibabaWdkWholesaleOrderCommitRequest() *AlibabaWdkWholesaleOrderCommitRequest{
    return &AlibabaWdkWholesaleOrderCommitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkWholesaleOrderCommitRequest) GetApiMethodName() string {
    return "alibaba.wdk.wholesale.order.commit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkWholesaleOrderCommitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCommitReq Setter
// 采购单信息
func (r *AlibabaWdkWholesaleOrderCommitRequest) SetOrderCommitReq(orderCommitReq *OrderCommitReq) error {
    r.orderCommitReq = orderCommitReq
    r.Set("order_commit_req", orderCommitReq)
    return nil
}

// OrderCommitReq Getter
func (r AlibabaWdkWholesaleOrderCommitRequest) GetOrderCommitReq() *OrderCommitReq {
    return r.orderCommitReq
}
