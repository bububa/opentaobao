package alihealthmdeer

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmdeer"
)

// AlibabaAlihealthMdeerScienceSynVideo 视频同步【保存/更新】
// alibaba.alihealth.mdeer.science.synVideo
//
// 视频同步【保存/更新】
func AlibabaAlihealthMdeerScienceSynVideo(ctx context.Context, clt *core.SDKClient, req *alihealthmdeer.AlibabaAlihealthMdeerScienceSynVideoAPIRequest, resp *alihealthmdeer.AlibabaAlihealthMdeerScienceSynVideoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
