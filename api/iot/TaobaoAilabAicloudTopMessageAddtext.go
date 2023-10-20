package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopMessageAddtext 精灵代说
// taobao.ailab.aicloud.top.message.addtext
//
// 精灵代说
func TaobaoAilabAicloudTopMessageAddtext(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopMessageAddtextAPIRequest, resp *iot.TaobaoAilabAicloudTopMessageAddtextAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
