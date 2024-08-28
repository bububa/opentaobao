package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// AlitripGrouptourProductUpload 新版跟团游商品维护接口
// alitrip.grouptour.product.upload
//
// 新版跟团游商品维护接口
func AlitripGrouptourProductUpload(ctx context.Context, clt *core.SDKClient, req *travel.AlitripGrouptourProductUploadAPIRequest, resp *travel.AlitripGrouptourProductUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
