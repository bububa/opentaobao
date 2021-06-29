package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口订单拉取 API请求
alibaba.wdk.order.list

五道口交易订单拉取接口
*/
type AlibabaWdkOrderListRequest struct {
    model.Params
    // 查询参数
    _batchQueryRequest   *BatchQueryRequest
}

// 初始化AlibabaWdkOrderListRequest对象
func NewAlibabaWdkOrderListRequest() *AlibabaWdkOrderListRequest{
    return &AlibabaWdkOrderListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderListRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BatchQueryRequest Setter
// 查询参数
func (r *AlibabaWdkOrderListRequest) SetBatchQueryRequest(_batchQueryRequest *BatchQueryRequest) error {
    r._batchQueryRequest = _batchQueryRequest
    r.Set("batch_query_request", _batchQueryRequest)
    return nil
}

// BatchQueryRequest Getter
func (r AlibabaWdkOrderListRequest) GetBatchQueryRequest() *BatchQueryRequest {
    return r._batchQueryRequest
}
