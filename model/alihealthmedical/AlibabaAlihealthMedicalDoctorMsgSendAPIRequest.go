package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicaldoctormsgsendAPIRequest 三方医生消息写入 API请求
// alibaba.alihealth.medical.doctor.msg.send
//
// 三方机构医生端发送消息同步写入阿里健康IM
type AlibabaalihealthmedicaldoctormsgsendAPIRequest struct {
	model.Params
	// request
	_inquiry *OuterMsgPullRequest
}

// NewAlibabaalihealthmedicaldoctormsgsendRequest 初始化AlibabaalihealthmedicaldoctormsgsendAPIRequest对象
func NewAlibabaalihealthmedicaldoctormsgsendRequest() *AlibabaalihealthmedicaldoctormsgsendAPIRequest {
	return &AlibabaalihealthmedicaldoctormsgsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicaldoctormsgsendAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.doctor.msg.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicaldoctormsgsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicaldoctormsgsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInquiry is Inquiry Setter
// request
func (r *AlibabaalihealthmedicaldoctormsgsendAPIRequest) SetInquiry(_inquiry *OuterMsgPullRequest) error {
	r._inquiry = _inquiry
	r.Set("inquiry", _inquiry)
	return nil
}

// GetInquiry Inquiry Getter
func (r AlibabaalihealthmedicaldoctormsgsendAPIRequest) GetInquiry() *OuterMsgPullRequest {
	return r._inquiry
}
