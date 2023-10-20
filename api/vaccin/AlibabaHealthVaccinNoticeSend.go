package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabahealthvaccinnoticesend 发送消息提醒
// alibaba.health.vaccin.notice.send
//
// ISV 通过免疫规划中心给用户发送短信或者支付宝 PUSH 提醒。
func Alibabahealthvaccinnoticesend(clt *core.SDKClient, req *vaccin.AlibabahealthvaccinnoticesendAPIRequest, session string) (*vaccin.AlibabahealthvaccinnoticesendAPIResponse, error) {
	var resp vaccin.AlibabahealthvaccinnoticesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
