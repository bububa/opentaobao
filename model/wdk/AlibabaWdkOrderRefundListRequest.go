package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口交易退款批量查询 API请求
alibaba.wdk.order.refund.list

按照条件查询退款数据
*/
type AlibabaWdkOrderRefundListRequest struct {
    model.Params
    // 查询条件
    _batchQueryRefundRequest   *BatchQueryRefundRequest
}

// 初始化AlibabaWdkOrderRefundListRequest对象
func NewAlibabaWdkOrderRefundListRequest() *AlibabaWdkOrderRefundListRequest{
    return &AlibabaWdkOrderRefundListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderRefundListRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.refund.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderRefundListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BatchQueryRefundRequest Setter
// 查询条件
func (r *AlibabaWdkOrderRefundListRequest) SetBatchQueryRefundRequest(_batchQueryRefundRequest *BatchQueryRefundRequest) error {
    r._batchQueryRefundRequest = _batchQueryRefundRequest
    r.Set("batch_query_refund_request", _batchQueryRefundRequest)
    return nil
}

// BatchQueryRefundRequest Getter
func (r AlibabaWdkOrderRefundListRequest) GetBatchQueryRefundRequest() *BatchQueryRefundRequest {
    return r._batchQueryRefundRequest
}
