package aedropshiper

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsImageSearch 图片搜索
// aliexpress.ds.image.search
//
// 图片搜索
func AliexpressDsImageSearch(ctx context.Context, clt *core.SDKClient, req *aedropshiper.AliexpressDsImageSearchAPIRequest, resp *aedropshiper.AliexpressDsImageSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
