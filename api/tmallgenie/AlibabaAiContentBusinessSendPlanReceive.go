package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaaicontentbusinesssendplanreceive 天猫精灵商业化采销发放计划领取
// alibaba.ai.content.business.send.plan.receive
//
// 天猫精灵商业化采销发放计划领取
func Alibabaaicontentbusinesssendplanreceive(clt *core.SDKClient, req *tmallgenie.AlibabaaicontentbusinesssendplanreceiveAPIRequest, session string) (*tmallgenie.AlibabaaicontentbusinesssendplanreceiveAPIResponse, error) {
	var resp tmallgenie.AlibabaaicontentbusinesssendplanreceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
