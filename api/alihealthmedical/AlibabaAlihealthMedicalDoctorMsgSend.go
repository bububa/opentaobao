package alihealthmedical

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmedical"
)

// AlibabaAlihealthMedicalDoctorMsgSend 三方医生消息写入
// alibaba.alihealth.medical.doctor.msg.send
//
// 三方机构医生端发送消息同步写入阿里健康IM
func AlibabaAlihealthMedicalDoctorMsgSend(ctx context.Context, clt *core.SDKClient, req *alihealthmedical.AlibabaAlihealthMedicalDoctorMsgSendAPIRequest, resp *alihealthmedical.AlibabaAlihealthMedicalDoctorMsgSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
