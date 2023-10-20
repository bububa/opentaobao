package vaccin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthVaccineRegisterSubmitAPIRequest cdc回传疫苗登记数据 API请求
// alibaba.alihealth.vaccine.register.submit
//
// cdc回传疫苗登记信息
type AlibabaAlihealthVaccineRegisterSubmitAPIRequest struct {
	model.Params
	// 无
	_topRequest *SubmitVcRegisterRequest
}

// NewAlibabaAlihealthVaccineRegisterSubmitRequest 初始化AlibabaAlihealthVaccineRegisterSubmitAPIRequest对象
func NewAlibabaAlihealthVaccineRegisterSubmitRequest() *AlibabaAlihealthVaccineRegisterSubmitAPIRequest {
	return &AlibabaAlihealthVaccineRegisterSubmitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthVaccineRegisterSubmitAPIRequest) Reset() {
	r._topRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthVaccineRegisterSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.vaccine.register.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthVaccineRegisterSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthVaccineRegisterSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopRequest is TopRequest Setter
// 无
func (r *AlibabaAlihealthVaccineRegisterSubmitAPIRequest) SetTopRequest(_topRequest *SubmitVcRegisterRequest) error {
	r._topRequest = _topRequest
	r.Set("top_request", _topRequest)
	return nil
}

// GetTopRequest TopRequest Getter
func (r AlibabaAlihealthVaccineRegisterSubmitAPIRequest) GetTopRequest() *SubmitVcRegisterRequest {
	return r._topRequest
}

var poolAlibabaAlihealthVaccineRegisterSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthVaccineRegisterSubmitRequest()
	},
}

// GetAlibabaAlihealthVaccineRegisterSubmitRequest 从 sync.Pool 获取 AlibabaAlihealthVaccineRegisterSubmitAPIRequest
func GetAlibabaAlihealthVaccineRegisterSubmitAPIRequest() *AlibabaAlihealthVaccineRegisterSubmitAPIRequest {
	return poolAlibabaAlihealthVaccineRegisterSubmitAPIRequest.Get().(*AlibabaAlihealthVaccineRegisterSubmitAPIRequest)
}

// ReleaseAlibabaAlihealthVaccineRegisterSubmitAPIRequest 将 AlibabaAlihealthVaccineRegisterSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthVaccineRegisterSubmitAPIRequest(v *AlibabaAlihealthVaccineRegisterSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthVaccineRegisterSubmitAPIRequest.Put(v)
}
