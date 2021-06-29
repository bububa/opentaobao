package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达退款单按门店聚合查询 API请求
alibaba.wdk.refund.aggregate

淘鲜达退款单按门店聚合查询
*/
type AlibabaWdkRefundAggregateRequest struct {
    model.Params
    // 系统自动生成
    _refundAggregateQueryRequest   *RefundAggregateQueryRequest
}

// 初始化AlibabaWdkRefundAggregateRequest对象
func NewAlibabaWdkRefundAggregateRequest() *AlibabaWdkRefundAggregateRequest{
    return &AlibabaWdkRefundAggregateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkRefundAggregateRequest) GetApiMethodName() string {
    return "alibaba.wdk.refund.aggregate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkRefundAggregateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundAggregateQueryRequest Setter
// 系统自动生成
func (r *AlibabaWdkRefundAggregateRequest) SetRefundAggregateQueryRequest(_refundAggregateQueryRequest *RefundAggregateQueryRequest) error {
    r._refundAggregateQueryRequest = _refundAggregateQueryRequest
    r.Set("refund_aggregate_query_request", _refundAggregateQueryRequest)
    return nil
}

// RefundAggregateQueryRequest Getter
func (r AlibabaWdkRefundAggregateRequest) GetRefundAggregateQueryRequest() *RefundAggregateQueryRequest {
    return r._refundAggregateQueryRequest
}
