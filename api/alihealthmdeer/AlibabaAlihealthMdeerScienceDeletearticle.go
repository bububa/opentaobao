package alihealthmdeer

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmdeer"
)

// AlibabaAlihealthMdeerScienceDeletearticle 文章删除
// alibaba.alihealth.mdeer.science.deletearticle
//
// 三方同步文章删除
func AlibabaAlihealthMdeerScienceDeletearticle(ctx context.Context, clt *core.SDKClient, req *alihealthmdeer.AlibabaAlihealthMdeerScienceDeletearticleAPIRequest, resp *alihealthmdeer.AlibabaAlihealthMdeerScienceDeletearticleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
