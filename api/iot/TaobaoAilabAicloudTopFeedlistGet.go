package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopFeedlistGet 获取对话流列表
// taobao.ailab.aicloud.top.feedlist.get
//
// 获取指定应用的对话流信息
func TaobaoAilabAicloudTopFeedlistGet(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopFeedlistGetAPIRequest, resp *iot.TaobaoAilabAicloudTopFeedlistGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
