package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscGrowthInteractiveSnsConverturl 防封接口
// alibaba.alsc.growth.interactive.sns.converturl
//
// 防封接口
func AlibabaAlscGrowthInteractiveSnsConverturl(clt *core.SDKClient, req *alsc.AlibabaAlscGrowthInteractiveSnsConverturlAPIRequest, session string) (*alsc.AlibabaAlscGrowthInteractiveSnsConverturlAPIResponse, error) {
	var resp alsc.AlibabaAlscGrowthInteractiveSnsConverturlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
