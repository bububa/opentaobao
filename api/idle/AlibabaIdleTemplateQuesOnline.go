package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTemplateQuesOnline 预上线SPU问卷版本
// alibaba.idle.template.ques.online
//
// 获取SPU最新版本问卷
func AlibabaIdleTemplateQuesOnline(clt *core.SDKClient, req *idle.AlibabaIdleTemplateQuesOnlineAPIRequest, resp *idle.AlibabaIdleTemplateQuesOnlineAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
