package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTemplateQuesGet 获取SPU最新版本问卷
// alibaba.idle.template.ques.get
//
// 获取SPU最新版本问卷
func AlibabaIdleTemplateQuesGet(clt *core.SDKClient, req *idle.AlibabaIdleTemplateQuesGetAPIRequest, resp *idle.AlibabaIdleTemplateQuesGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
