package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaFcMallxInteractionAiPicList 花园ai作画定制查询
// alibaba.fc.mallx.interaction.ai.pic.list
//
// 花园ai作画定制查询
func AlibabaFcMallxInteractionAiPicList(clt *core.SDKClient, req *interact.AlibabaFcMallxInteractionAiPicListAPIRequest, resp *interact.AlibabaFcMallxInteractionAiPicListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
