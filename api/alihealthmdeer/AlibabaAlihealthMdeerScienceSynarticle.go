package alihealthmdeer

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmdeer"
)

// Alibabaalihealthmdeersciencesynarticle 医知鹿文章同步【保存/更新】
// alibaba.alihealth.mdeer.science.synarticle
//
// 文章同步【保存/更新】
func Alibabaalihealthmdeersciencesynarticle(clt *core.SDKClient, req *alihealthmdeer.AlibabaalihealthmdeersciencesynarticleAPIRequest, session string) (*alihealthmdeer.AlibabaalihealthmdeersciencesynarticleAPIResponse, error) {
	var resp alihealthmdeer.AlibabaalihealthmdeersciencesynarticleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
