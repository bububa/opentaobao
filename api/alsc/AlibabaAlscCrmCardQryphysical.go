package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCardQryphysical 查询物理卡
// alibaba.alsc.crm.card.qryphysical
//
// 查询物理卡
func AlibabaAlscCrmCardQryphysical(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardQryphysicalAPIRequest, resp *alsc.AlibabaAlscCrmCardQryphysicalAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
