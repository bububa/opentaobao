package mtopopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// AlibabaInteractAllsparkisvDraw allspark提供抽奖tida接口对应鉴权接口
// alibaba.interact.allsparkisv.draw
//
// 该接口没有实际对外使用。只是内部鉴权使用，不会有三方应用调用
func AlibabaInteractAllsparkisvDraw(ctx context.Context, clt *core.SDKClient, req *mtopopen.AlibabaInteractAllsparkisvDrawAPIRequest, resp *mtopopen.AlibabaInteractAllsparkisvDrawAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
