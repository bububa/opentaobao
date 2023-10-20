package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// Aliexpressdsimagesearch 图片搜索
// aliexpress.ds.image.search
//
// 图片搜索
func Aliexpressdsimagesearch(clt *core.SDKClient, req *aedropshiper.AliexpressdsimagesearchAPIRequest, session string) (*aedropshiper.AliexpressdsimagesearchAPIResponse, error) {
	var resp aedropshiper.AliexpressdsimagesearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
