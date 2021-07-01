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
type AlibabaWdkWholesaleOrderCommitAPIRequest struct {
    model.Params
    // 采购单信息
    _orderCommitReq   *OrderCommitReq
}

// 初始化AlibabaWdkWholesaleOrderCommitAPIRequest对象
func NewAlibabaWdkWholesaleOrderCommitRequest() *AlibabaWdkWholesaleOrderCommitAPIRequest{
    return &AlibabaWdkWholesaleOrderCommitAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkWholesaleOrderCommitAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.wholesale.order.commit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkWholesaleOrderCommitAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCommitReq Setter
// 采购单信息
func (r *AlibabaWdkWholesaleOrderCommitAPIRequest) SetOrderCommitReq(_orderCommitReq *OrderCommitReq) error {
    r._orderCommitReq = _orderCommitReq
    r.Set("order_commit_req", _orderCommitReq)
    return nil
}

// OrderCommitReq Getter
func (r AlibabaWdkWholesaleOrderCommitAPIRequest) GetOrderCommitReq() *OrderCommitReq {
    return r._orderCommitReq
}
