package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// Aliexpressaffiliateimagesearch 图搜
// aliexpress.affiliate.image.search
//
// 图片搜索接口
func Aliexpressaffiliateimagesearch(clt *core.SDKClient, req *aecreatives.AliexpressaffiliateimagesearchAPIRequest, session string) (*aecreatives.AliexpressaffiliateimagesearchAPIResponse, error) {
	var resp aecreatives.AliexpressaffiliateimagesearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
