package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscGrowthInteractiveSnsConverturl 防封接口
// alibaba.alsc.growth.interactive.sns.converturl
//
// 防封接口
func AlibabaAlscGrowthInteractiveSnsConverturl(clt *core.SDKClient, req *alsc.AlibabaAlscGrowthInteractiveSnsConverturlAPIRequest, resp *alsc.AlibabaAlscGrowthInteractiveSnsConverturlAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
