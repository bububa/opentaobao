package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCardBatchActive 批量激活卡
// alibaba.alsc.crm.card.batch.active
//
// 批量激活卡
func AlibabaAlscCrmCardBatchActive(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardBatchActiveAPIRequest, resp *alsc.AlibabaAlscCrmCardBatchActiveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
