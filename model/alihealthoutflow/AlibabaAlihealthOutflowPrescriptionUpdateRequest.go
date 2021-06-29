package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-修改处方 API请求
alibaba.alihealth.outflow.prescription.update

阿里健康-处方外流-对外提供处方修改功能
*/
type AlibabaAlihealthOutflowPrescriptionUpdateRequest struct {
    model.Params
    // 入参对象
    _updateRequest   *PrescriptionOutflowUpdateRequest
}

// 初始化AlibabaAlihealthOutflowPrescriptionUpdateRequest对象
func NewAlibabaAlihealthOutflowPrescriptionUpdateRequest() *AlibabaAlihealthOutflowPrescriptionUpdateRequest{
    return &AlibabaAlihealthOutflowPrescriptionUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowPrescriptionUpdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.prescription.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowPrescriptionUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UpdateRequest Setter
// 入参对象
func (r *AlibabaAlihealthOutflowPrescriptionUpdateRequest) SetUpdateRequest(_updateRequest *PrescriptionOutflowUpdateRequest) error {
    r._updateRequest = _updateRequest
    r.Set("update_request", _updateRequest)
    return nil
}

// UpdateRequest Getter
func (r AlibabaAlihealthOutflowPrescriptionUpdateRequest) GetUpdateRequest() *PrescriptionOutflowUpdateRequest {
    return r._updateRequest
}
