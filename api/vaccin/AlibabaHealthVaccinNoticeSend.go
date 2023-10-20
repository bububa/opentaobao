package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinNoticeSend 发送消息提醒
// alibaba.health.vaccin.notice.send
//
// ISV 通过免疫规划中心给用户发送短信或者支付宝 PUSH 提醒。
func AlibabaHealthVaccinNoticeSend(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeSendAPIRequest, resp *vaccin.AlibabaHealthVaccinNoticeSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
