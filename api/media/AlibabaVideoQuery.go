package media

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// AlibabaVideoQuery 查询视频信息
// alibaba.video.query
//
// 查询视频信息
func AlibabaVideoQuery(ctx context.Context, clt *core.SDKClient, req *media.AlibabaVideoQueryAPIRequest, resp *media.AlibabaVideoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
