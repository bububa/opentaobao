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
type AlibabaAlihealthOutflowPrescriptionCreateAPIRequest struct {
    model.Params
    // 保存处方入参
    _createRequest   *PrescriptionOutflowUpdateRequest
}

// 初始化AlibabaAlihealthOutflowPrescriptionCreateAPIRequest对象
func NewAlibabaAlihealthOutflowPrescriptionCreateRequest() *AlibabaAlihealthOutflowPrescriptionCreateAPIRequest{
    return &AlibabaAlihealthOutflowPrescriptionCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowPrescriptionCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.prescription.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowPrescriptionCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreateRequest Setter
// 保存处方入参
func (r *AlibabaAlihealthOutflowPrescriptionCreateAPIRequest) SetCreateRequest(_createRequest *PrescriptionOutflowUpdateRequest) error {
    r._createRequest = _createRequest
    r.Set("create_request", _createRequest)
    return nil
}

// CreateRequest Getter
func (r AlibabaAlihealthOutflowPrescriptionCreateAPIRequest) GetCreateRequest() *PrescriptionOutflowUpdateRequest {
    return r._createRequest
}
