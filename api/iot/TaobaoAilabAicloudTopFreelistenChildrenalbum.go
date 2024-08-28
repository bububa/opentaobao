package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopFreelistenChildrenalbum 儿童音频列表
// taobao.ailab.aicloud.top.freelisten.childrenalbum
//
// 儿童音频列表
func TaobaoAilabAicloudTopFreelistenChildrenalbum(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest, resp *iot.TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
