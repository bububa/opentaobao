package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCardQryphysical 查询物理卡
// alibaba.alsc.crm.card.qryphysical
//
// 查询物理卡
func AlibabaAlscCrmCardQryphysical(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardQryphysicalAPIRequest, resp *alsc.AlibabaAlscCrmCardQryphysicalAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
