package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-创建处方 API请求
alibaba.alihealth.outflow.prescription.create

阿里健康-处方外流-对外提供保存处方功能
*/
type AlibabaAlihealthOutflowPrescriptionCreateRequest struct {
    model.Params
    // 保存处方入参
    _createRequest   *PrescriptionOutflowUpdateRequest
}

// 初始化AlibabaAlihealthOutflowPrescriptionCreateRequest对象
func NewAlibabaAlihealthOutflowPrescriptionCreateRequest() *AlibabaAlihealthOutflowPrescriptionCreateRequest{
    return &AlibabaAlihealthOutflowPrescriptionCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowPrescriptionCreateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.prescription.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowPrescriptionCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreateRequest Setter
// 保存处方入参
func (r *AlibabaAlihealthOutflowPrescriptionCreateRequest) SetCreateRequest(_createRequest *PrescriptionOutflowUpdateRequest) error {
    r._createRequest = _createRequest
    r.Set("create_request", _createRequest)
    return nil
}

// CreateRequest Getter
func (r AlibabaAlihealthOutflowPrescriptionCreateRequest) GetCreateRequest() *PrescriptionOutflowUpdateRequest {
    return r._createRequest
}
