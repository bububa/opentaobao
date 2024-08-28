package media

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// AlibabaVideoTokenGet 获取上传token
// alibaba.video.token.get
//
// 获取上传token
func AlibabaVideoTokenGet(ctx context.Context, clt *core.SDKClient, req *media.AlibabaVideoTokenGetAPIRequest, resp *media.AlibabaVideoTokenGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
