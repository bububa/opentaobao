package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopLikeDelete 取消收藏
// taobao.ailab.aicloud.top.like.delete
//
// 取消收藏
func TaobaoAilabAicloudTopLikeDelete(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopLikeDeleteAPIRequest, resp *iot.TaobaoAilabAicloudTopLikeDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
