package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaaicontentbusinessgetthirdcyclevipstatus 天猫精灵商业化获取三方连续包会员状态
// alibaba.ai.content.business.get.third.cycle.vip.status
//
// 天猫精灵商业化获取三方连续包会员状态
func Alibabaaicontentbusinessgetthirdcyclevipstatus(clt *core.SDKClient, req *tmallgenie.AlibabaaicontentbusinessgetthirdcyclevipstatusAPIRequest, session string) (*tmallgenie.AlibabaaicontentbusinessgetthirdcyclevipstatusAPIResponse, error) {
	var resp tmallgenie.AlibabaaicontentbusinessgetthirdcyclevipstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
