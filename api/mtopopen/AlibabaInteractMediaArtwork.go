package mtopopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// AlibabaInteractMediaArtwork 原图相关鉴权接口
// alibaba.interact.media.artwork
//
// 拍摄并上传原图相关鉴权接口
func AlibabaInteractMediaArtwork(ctx context.Context, clt *core.SDKClient, req *mtopopen.AlibabaInteractMediaArtworkAPIRequest, resp *mtopopen.AlibabaInteractMediaArtworkAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
