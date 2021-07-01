package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalOrderRefuseAPIRequest
三方机构通知平台"医生拒诊" API请求
alibaba.alihealth.medical.order.refuse

三方机构通知平台"医生拒诊" */
type AlibabaAlihealthMedicalOrderRefuseAPIRequest struct {
	model.Params
	// 请求入参
	_requestInfo *RefuseOrderRequestDto
}

// NewAlibabaAlihealthMedicalOrderRefuseRequest 初始化AlibabaAlihealthMedicalOrderRefuseAPIRequest对象
func NewAlibabaAlihealthMedicalOrderRefuseRequest() *AlibabaAlihealthMedicalOrderRefuseAPIRequest {
	return &AlibabaAlihealthMedicalOrderRefuseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalOrderRefuseAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.order.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalOrderRefuseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RequestInfo Setter
// 请求入参
func (r *AlibabaAlihealthMedicalOrderRefuseAPIRequest) SetRequestInfo(_requestInfo *RefuseOrderRequestDto) error {
	r._requestInfo = _requestInfo
	r.Set("request_info", _requestInfo)
	return nil
}

// Get RequestInfo Getter
func (r AlibabaAlihealthMedicalOrderRefuseAPIRequest) GetRequestInfo() *RefuseOrderRequestDto {
	return r._requestInfo
}
