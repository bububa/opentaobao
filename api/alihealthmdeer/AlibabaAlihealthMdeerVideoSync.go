package alihealthmdeer

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmdeer"
)

// AlibabaAlihealthMdeerVideoSync 合作伙伴视频同步给医知鹿接口
// alibaba.alihealth.mdeer.video.sync
//
// 合伙做伴内容同步接口，用来视频同步
func AlibabaAlihealthMdeerVideoSync(ctx context.Context, clt *core.SDKClient, req *alihealthmdeer.AlibabaAlihealthMdeerVideoSyncAPIRequest, resp *alihealthmdeer.AlibabaAlihealthMdeerVideoSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
