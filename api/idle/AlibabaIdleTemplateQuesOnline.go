package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidletemplatequesonline 预上线SPU问卷版本
// alibaba.idle.template.ques.online
//
// 获取SPU最新版本问卷
func Alibabaidletemplatequesonline(clt *core.SDKClient, req *idle.AlibabaidletemplatequesonlineAPIRequest, session string) (*idle.AlibabaidletemplatequesonlineAPIResponse, error) {
	var resp idle.AlibabaidletemplatequesonlineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
