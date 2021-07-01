package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

/* AlibabaHealthVaccinNoticeSend
发送消息提醒
alibaba.health.vaccin.notice.send

ISV 通过免疫规划中心给用户发送短信或者支付宝 PUSH 提醒。 */
func AlibabaHealthVaccinNoticeSend(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeSendAPIRequest, session string) (*vaccin.AlibabaHealthVaccinNoticeSendAPIResponse, error) {
	var resp vaccin.AlibabaHealthVaccinNoticeSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
