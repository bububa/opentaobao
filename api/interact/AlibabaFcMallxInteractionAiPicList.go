package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaFcMallxInteractionAiPicList 花园ai作画定制查询
// alibaba.fc.mallx.interaction.ai.pic.list
//
// 花园ai作画定制查询
func AlibabaFcMallxInteractionAiPicList(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaFcMallxInteractionAiPicListAPIRequest, resp *interact.AlibabaFcMallxInteractionAiPicListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
