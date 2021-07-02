package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalRegisterWeiyiSyncAPIRequest 微医数据号源回传 API请求
// alibaba.alihealth.medical.register.weiyi.sync
//
// 微医号源数据回传
type AlibabaAlihealthMedicalRegisterWeiyiSyncAPIRequest struct {
	model.Params
	// 号源数据实体
	_serviceRequest *SourcesReturnVo
}

// NewAlibabaAlihealthMedicalRegisterWeiyiSyncRequest 初始化AlibabaAlihealthMedicalRegisterWeiyiSyncAPIRequest对象
func NewAlibabaAlihealthMedicalRegisterWeiyiSyncRequest() *AlibabaAlihealthMedicalRegisterWeiyiSyncAPIRequest {
	return &AlibabaAlihealthMedicalRegisterWeiyiSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalRegisterWeiyiSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.register.weiyi.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalRegisterWeiyiSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ServiceRequest Setter
// 号源数据实体
func (r *AlibabaAlihealthMedicalRegisterWeiyiSyncAPIRequest) SetServiceRequest(_serviceRequest *SourcesReturnVo) error {
	r._serviceRequest = _serviceRequest
	r.Set("service_request", _serviceRequest)
	return nil
}

// Get ServiceRequest Getter
func (r AlibabaAlihealthMedicalRegisterWeiyiSyncAPIRequest) GetServiceRequest() *SourcesReturnVo {
	return r._serviceRequest
}
