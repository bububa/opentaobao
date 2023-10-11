package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAiContentBusinessGetThirdCycleVipStatus 天猫精灵商业化获取三方连续包会员状态
// alibaba.ai.content.business.get.third.cycle.vip.status
//
// 天猫精灵商业化获取三方连续包会员状态
func AlibabaAiContentBusinessGetThirdCycleVipStatus(clt *core.SDKClient, req *tmallgenie.AlibabaAiContentBusinessGetThirdCycleVipStatusAPIRequest, session string) (*tmallgenie.AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse, error) {
	var resp tmallgenie.AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
