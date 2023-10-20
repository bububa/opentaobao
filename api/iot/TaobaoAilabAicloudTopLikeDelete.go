package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopLikeDelete 取消收藏
// taobao.ailab.aicloud.top.like.delete
//
// 取消收藏
func TaobaoAilabAicloudTopLikeDelete(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopLikeDeleteAPIRequest, resp *iot.TaobaoAilabAicloudTopLikeDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
