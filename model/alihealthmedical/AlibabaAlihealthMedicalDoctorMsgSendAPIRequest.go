package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalDoctorMsgSendAPIRequest
三方医生消息写入 API请求
alibaba.alihealth.medical.doctor.msg.send

三方机构医生端发送消息同步写入阿里健康IM */
type AlibabaAlihealthMedicalDoctorMsgSendAPIRequest struct {
	model.Params
	// request
	_inquiry *OuterMsgPullRequest
}

// New
