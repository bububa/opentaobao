package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopLikeAdd 增加收藏
// taobao.ailab.aicloud.top.like.add
//
// 将制定内容加入收藏
func TaobaoAilabAicloudTopLikeAdd(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopLikeAddAPIRequest, resp *iot.TaobaoAilabAicloudTopLikeAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
