package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopLikeAdd 增加收藏
// taobao.ailab.aicloud.top.like.add
//
// 将制定内容加入收藏
func TaobaoAilabAicloudTopLikeAdd(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopLikeAddAPIRequest, resp *iot.TaobaoAilabAicloudTopLikeAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
