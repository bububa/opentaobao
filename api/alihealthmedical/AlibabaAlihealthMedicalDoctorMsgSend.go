package alihealthmedical

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmedical"
)

/* AlibabaAlihealthMedicalDoctorMsgSend
三方医生消息写入
alibaba.alihealth.medical.doctor.msg.send

三方机构医生端发送消息同步写入阿里健康IM */
func AlibabaAlihealthMedicalDoctorMsgSend(clt *core.SDKClient, req *alihealthmedical.AlibabaAlihealthMedicalDoctorMsgSendAPIRequest, session string) (*alihealthmedical.AlibabaAlihealthMedicalDoctorMsgSendAPIResponse, error) {
	var resp alihealthmedical.AlibabaAlihealthMedicalDoctorMsgSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
