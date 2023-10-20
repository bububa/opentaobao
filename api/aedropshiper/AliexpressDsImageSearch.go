package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsImageSearch 图片搜索
// aliexpress.ds.image.search
//
// 图片搜索
func AliexpressDsImageSearch(clt *core.SDKClient, req *aedropshiper.AliexpressDsImageSearchAPIRequest, session string) (*aedropshiper.AliexpressDsImageSearchAPIResponse, error) {
	var resp aedropshiper.AliexpressDsImageSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
