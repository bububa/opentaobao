package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// Alitripgrouptourproductupload 新版跟团游商品维护接口
// alitrip.grouptour.product.upload
//
// 新版跟团游商品维护接口
func Alitripgrouptourproductupload(clt *core.SDKClient, req *travel.AlitripgrouptourproductuploadAPIRequest, session string) (*travel.AlitripgrouptourproductuploadAPIResponse, error) {
	var resp travel.AlitripgrouptourproductuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
