package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCardActive 标准激活卡
// alibaba.alsc.crm.card.active
//
// 激活卡
func AlibabaAlscCrmCardActive(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardActiveAPIRequest, resp *alsc.AlibabaAlscCrmCardActiveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
