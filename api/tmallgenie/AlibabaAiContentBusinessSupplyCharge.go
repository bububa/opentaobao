package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaaicontentbusinesssupplycharge 供销商品充值接口
// alibaba.ai.content.business.supply.charge
//
// 供销商品充值接口
func Alibabaaicontentbusinesssupplycharge(clt *core.SDKClient, req *tmallgenie.AlibabaaicontentbusinesssupplychargeAPIRequest, session string) (*tmallgenie.AlibabaaicontentbusinesssupplychargeAPIResponse, error) {
	var resp tmallgenie.AlibabaaicontentbusinesssupplychargeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
