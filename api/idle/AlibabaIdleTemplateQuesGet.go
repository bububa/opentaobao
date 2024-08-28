package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTemplateQuesGet 获取SPU最新版本问卷
// alibaba.idle.template.ques.get
//
// 获取SPU最新版本问卷
func AlibabaIdleTemplateQuesGet(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleTemplateQuesGetAPIRequest, resp *idle.AlibabaIdleTemplateQuesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
