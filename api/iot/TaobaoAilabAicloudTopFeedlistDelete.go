package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopFeedlistDelete 删除单条对话流信息
// taobao.ailab.aicloud.top.feedlist.delete
//
// 删除指定的某一条对话流信息
func TaobaoAilabAicloudTopFeedlistDelete(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopFeedlistDeleteAPIRequest, resp *iot.TaobaoAilabAicloudTopFeedlistDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
