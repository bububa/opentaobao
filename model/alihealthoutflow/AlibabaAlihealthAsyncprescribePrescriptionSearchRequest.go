package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
异步开方处方查询 API请求
alibaba.alihealth.asyncprescribe.prescription.search

异步开方处方查询
*/
type AlibabaAlihealthAsyncprescribePrescriptionSearchRequest struct {
    model.Params
    // 查询入参
    _searchRequest   *AsyncPrescribeSearchRequest
}

// 初始化AlibabaAlihealthAsyncprescribePrescriptionSearchRequest对象
func NewAlibabaAlihealthAsyncprescribePrescriptionSearchRequest() *AlibabaAlihealthAsyncprescribePrescriptionSearchRequest{
    return &AlibabaAlihealthAsyncprescribePrescriptionSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAsyncprescribePrescriptionSearchRequest) GetApiMethodName() string {
    return "alibaba.alihealth.asyncprescribe.prescription.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAsyncprescribePrescriptionSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SearchRequest Setter
// 查询入参
func (r *AlibabaAlihealthAsyncprescribePrescriptionSearchRequest) SetSearchRequest(_searchRequest *AsyncPrescribeSearchRequest) error {
    r._searchRequest = _searchRequest
    r.Set("search_request", _searchRequest)
    return nil
}

// SearchRequest Getter
func (r AlibabaAlihealthAsyncprescribePrescriptionSearchRequest) GetSearchRequest() *AsyncPrescribeSearchRequest {
    return r._searchRequest
}
