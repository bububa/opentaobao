package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAiContentBusinessSupplyCharge 供销商品充值接口
// alibaba.ai.content.business.supply.charge
//
// 供销商品充值接口
func AlibabaAiContentBusinessSupplyCharge(clt *core.SDKClient, req *tmallgenie.AlibabaAiContentBusinessSupplyChargeAPIRequest, session string) (*tmallgenie.AlibabaAiContentBusinessSupplyChargeAPIResponse, error) {
	var resp tmallgenie.AlibabaAiContentBusinessSupplyChargeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
