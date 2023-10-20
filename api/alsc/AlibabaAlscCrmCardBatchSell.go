package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCardBatchSell 批量开卡（售卡）
// alibaba.alsc.crm.card.batch.sell
//
// 批量开卡（售卡）
func AlibabaAlscCrmCardBatchSell(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardBatchSellAPIRequest, resp *alsc.AlibabaAlscCrmCardBatchSellAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
