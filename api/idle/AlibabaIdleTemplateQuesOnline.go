package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTemplateQuesOnline 预上线SPU问卷版本
// alibaba.idle.template.ques.online
//
// 获取SPU最新版本问卷
func AlibabaIdleTemplateQuesOnline(clt *core.SDKClient, req *idle.AlibabaIdleTemplateQuesOnlineAPIRequest, session string) (*idle.AlibabaIdleTemplateQuesOnlineAPIResponse, error) {
	var resp idle.AlibabaIdleTemplateQuesOnlineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
