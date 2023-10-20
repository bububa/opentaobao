package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// AlitripGrouptourProductUpload 新版跟团游商品维护接口
// alitrip.grouptour.product.upload
//
// 新版跟团游商品维护接口
func AlitripGrouptourProductUpload(clt *core.SDKClient, req *travel.AlitripGrouptourProductUploadAPIRequest, resp *travel.AlitripGrouptourProductUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
