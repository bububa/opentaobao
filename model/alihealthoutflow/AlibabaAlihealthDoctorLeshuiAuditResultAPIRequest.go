package alihealthoutflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest 乐税审核结果通知 API请求
// alibaba.alihealth.doctor.leshui.audit.result
//
// 乐税审核结果通知
type AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest struct {
	model.Params
	// 入参
	_param *PayTaxNoticeRequest
}

// NewAlibabaAlihealthDoctorLeshuiAuditResultRequest 初始化AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest对象
func NewAlibabaAlihealthDoctorLeshuiAuditResultRequest() *AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest {
	return &AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.doctor.leshui.audit.result"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest) SetParam(_param *PayTaxNoticeRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest) GetParam() *PayTaxNoticeRequest {
	return r._param
}

var poolAlibabaAlihealthDoctorLeshuiAuditResultAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDoctorLeshuiAuditResultRequest()
	},
}

// GetAlibabaAlihealthDoctorLeshuiAuditResultRequest 从 sync.Pool 获取 AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest
func GetAlibabaAlihealthDoctorLeshuiAuditResultAPIRequest() *AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest {
	return poolAlibabaAlihealthDoctorLeshuiAuditResultAPIRequest.Get().(*AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest)
}

// ReleaseAlibabaAlihealthDoctorLeshuiAuditResultAPIRequest 将 AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDoctorLeshuiAuditResultAPIRequest(v *AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDoctorLeshuiAuditResultAPIRequest.Put(v)
}
