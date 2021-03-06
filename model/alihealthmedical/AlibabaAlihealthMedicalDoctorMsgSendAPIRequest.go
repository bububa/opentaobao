package alihealthmedical

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalDoctorMsgSendAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.doctor.msg.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalDoctorMsgSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
