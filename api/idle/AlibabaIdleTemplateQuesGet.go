package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidletemplatequesget 获取SPU最新版本问卷
// alibaba.idle.template.ques.get
//
// 获取SPU最新版本问卷
func Alibabaidletemplatequesget(clt *core.SDKClient, req *idle.AlibabaidletemplatequesgetAPIRequest, session string) (*idle.AlibabaidletemplatequesgetAPIResponse, error) {
	var resp idle.AlibabaidletemplatequesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
