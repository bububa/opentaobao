package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

/* AlitripGrouptourProductUpload
新版跟团游商品维护接口
alitrip.grouptour.product.upload

新版跟团游商品维护接口 */
func AlitripGrouptourProductUpload(clt *core.SDKClient, req *travel.AlitripGrouptourProductUploadAPIRequest, session string) (*travel.AlitripGrouptourProductUploadAPIResponse, error) {
	var resp travel.AlitripGrouptourProductUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
