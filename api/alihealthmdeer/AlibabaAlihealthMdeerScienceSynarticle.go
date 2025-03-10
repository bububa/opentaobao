package alihealthmdeer

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmdeer"
)

// AlibabaAlihealthMdeerScienceSynarticle 医知鹿文章同步【保存/更新】
// alibaba.alihealth.mdeer.science.synarticle
//
// 文章同步【保存/更新】
func AlibabaAlihealthMdeerScienceSynarticle(ctx context.Context, clt *core.SDKClient, req *alihealthmdeer.AlibabaAlihealthMdeerScienceSynarticleAPIRequest, resp *alihealthmdeer.AlibabaAlihealthMdeerScienceSynarticleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
