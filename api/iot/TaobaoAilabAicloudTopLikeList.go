package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopLikeList 列出收藏列表
// taobao.ailab.aicloud.top.like.list
//
// 列出收藏列表
func TaobaoAilabAicloudTopLikeList(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopLikeListAPIRequest, resp *iot.TaobaoAilabAicloudTopLikeListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
