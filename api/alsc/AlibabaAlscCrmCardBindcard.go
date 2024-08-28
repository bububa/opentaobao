package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCardBindcard 绑定物理卡
// alibaba.alsc.crm.card.bindcard
//
// 将会员卡和实例物理卡绑定在一起
func AlibabaAlscCrmCardBindcard(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardBindcardAPIRequest, resp *alsc.AlibabaAlscCrmCardBindcardAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
