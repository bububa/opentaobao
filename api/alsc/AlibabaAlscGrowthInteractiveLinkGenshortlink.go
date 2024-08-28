package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscGrowthInteractiveLinkGenshortlink 短链接口
// alibaba.alsc.growth.interactive.link.genshortlink
//
// 短链接口
func AlibabaAlscGrowthInteractiveLinkGenshortlink(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest, resp *alsc.AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
