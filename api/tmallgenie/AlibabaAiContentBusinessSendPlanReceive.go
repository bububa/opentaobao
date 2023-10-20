package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAiContentBusinessSendPlanReceive 天猫精灵商业化采销发放计划领取
// alibaba.ai.content.business.send.plan.receive
//
// 天猫精灵商业化采销发放计划领取
func AlibabaAiContentBusinessSendPlanReceive(clt *core.SDKClient, req *tmallgenie.AlibabaAiContentBusinessSendPlanReceiveAPIRequest, session string) (*tmallgenie.AlibabaAiContentBusinessSendPlanReceiveAPIResponse, error) {
	var resp tmallgenie.AlibabaAiContentBusinessSendPlanReceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
