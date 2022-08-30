package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAiContentBusinessSendPlanQuery 内容商业化发放权益查询
// alibaba.ai.content.business.send.plan.query
//
// 天猫精灵内容生态的权益查询
func AlibabaAiContentBusinessSendPlanQuery(clt *core.SDKClient, req *tmallgenie.AlibabaAiContentBusinessSendPlanQueryAPIRequest, session string) (*tmallgenie.AlibabaAiContentBusinessSendPlanQueryAPIResponse, error) {
	var resp tmallgenie.AlibabaAiContentBusinessSendPlanQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
