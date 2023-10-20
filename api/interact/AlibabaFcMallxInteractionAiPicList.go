package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaFcMallxInteractionAiPicList 花园ai作画定制查询
// alibaba.fc.mallx.interaction.ai.pic.list
//
// 花园ai作画定制查询
func AlibabaFcMallxInteractionAiPicList(clt *core.SDKClient, req *interact.AlibabaFcMallxInteractionAiPicListAPIRequest, session string) (*interact.AlibabaFcMallxInteractionAiPicListAPIResponse, error) {
	var resp interact.AlibabaFcMallxInteractionAiPicListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
