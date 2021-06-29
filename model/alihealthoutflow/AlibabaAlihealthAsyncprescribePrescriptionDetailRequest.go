package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
异步开方处方详情 API请求
alibaba.alihealth.asyncprescribe.prescription.detail

异步开方处方查询
*/
type AlibabaAlihealthAsyncprescribePrescriptionDetailRequest struct {
    model.Params
    // 入参
    _detailRequest   *AsyncPrescribeDetailRequest
}

// 初始化AlibabaAlihealthAsyncprescribePrescriptionDetailRequest对象
func NewAlibabaAlihealthAsyncprescribePrescriptionDetailRequest() *AlibabaAlihealthAsyncprescribePrescriptionDetailRequest{
    return &AlibabaAlihealthAsyncprescribePrescriptionDetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAsyncprescribePrescriptionDetailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.asyncprescribe.prescription.detail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAsyncprescribePrescriptionDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DetailRequest Setter
// 入参
func (r *AlibabaAlihealthAsyncprescribePrescriptionDetailRequest) SetDetailRequest(_detailRequest *AsyncPrescribeDetailRequest) error {
    r._detailRequest = _detailRequest
    r.Set("detail_request", _detailRequest)
    return nil
}

// DetailRequest Getter
func (r AlibabaAlihealthAsyncprescribePrescriptionDetailRequest) GetDetailRequest() *AsyncPrescribeDetailRequest {
    return r._detailRequest
}
