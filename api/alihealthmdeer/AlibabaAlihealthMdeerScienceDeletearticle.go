package alihealthmdeer

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmdeer"
)

// Alibabaalihealthmdeersciencedeletearticle 文章删除
// alibaba.alihealth.mdeer.science.deletearticle
//
// 三方同步文章删除
func Alibabaalihealthmdeersciencedeletearticle(clt *core.SDKClient, req *alihealthmdeer.AlibabaalihealthmdeersciencedeletearticleAPIRequest, session string) (*alihealthmdeer.AlibabaalihealthmdeersciencedeletearticleAPIResponse, error) {
	var resp alihealthmdeer.AlibabaalihealthmdeersciencedeletearticleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
