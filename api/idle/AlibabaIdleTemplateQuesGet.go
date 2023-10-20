package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTemplateQuesGet 获取SPU最新版本问卷
// alibaba.idle.template.ques.get
//
// 获取SPU最新版本问卷
func AlibabaIdleTemplateQuesGet(clt *core.SDKClient, req *idle.AlibabaIdleTemplateQuesGetAPIRequest, session string) (*idle.AlibabaIdleTemplateQuesGetAPIResponse, error) {
	var resp idle.AlibabaIdleTemplateQuesGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
