package alihealthmedical

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalDoctorMsgSendAPIRequest 三方医生消息写入 API请求
// alibaba.alihealth.medical.doctor.msg.send
//
// 三方机构医生端发送消息同步写入阿里健康IM
type AlibabaAlihealthMedicalDoctorMsgSendAPIRequest struct {
	model.Params
	// request
	_inquiry *OuterMsgPullRequest
}

// NewAlibabaAlihealthMedicalDoctorMsgSendRequest 初始化AlibabaAlihealthMedicalDoctorMsgSendAPIRequest对象
func NewAlibabaAlihealthMedicalDoctorMsgSendRequest() *AlibabaAlihealthMedicalDoctorMsgSendAPIRequest {
	return &AlibabaAlihealthMedicalDoctorMsgSendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalDoctorMsgSendAPIRequest) Reset() {
	r._inquiry = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalDoctorMsgSendAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.doctor.msg.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalDoctorMsgSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalDoctorMsgSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInquiry is Inquiry Setter
// request
func (r *AlibabaAlihealthMedicalDoctorMsgSendAPIRequest) SetInquiry(_inquiry *OuterMsgPullRequest) error {
	r._inquiry = _inquiry
	r.Set("inquiry", _inquiry)
	return nil
}

// GetInquiry Inquiry Getter
func (r AlibabaAlihealthMedicalDoctorMsgSendAPIRequest) GetInquiry() *OuterMsgPullRequest {
	return r._inquiry
}

var poolAlibabaAlihealthMedicalDoctorMsgSendAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalDoctorMsgSendRequest()
	},
}

// GetAlibabaAlihealthMedicalDoctorMsgSendRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalDoctorMsgSendAPIRequest
func GetAlibabaAlihealthMedicalDoctorMsgSendAPIRequest() *AlibabaAlihealthMedicalDoctorMsgSendAPIRequest {
	return poolAlibabaAlihealthMedicalDoctorMsgSendAPIRequest.Get().(*AlibabaAlihealthMedicalDoctorMsgSendAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalDoctorMsgSendAPIRequest 将 AlibabaAlihealthMedicalDoctorMsgSendAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalDoctorMsgSendAPIRequest(v *AlibabaAlihealthMedicalDoctorMsgSendAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalDoctorMsgSendAPIRequest.Put(v)
}
