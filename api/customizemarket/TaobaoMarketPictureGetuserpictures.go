package customizemarket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/customizemarket"
)

// TaobaoMarketPictureGetuserpictures 读取用户上传图片
// taobao.market.picture.getuserpictures
//
// 商家通过用户信息，获取用户上传的
func TaobaoMarketPictureGetuserpictures(ctx context.Context, clt *core.SDKClient, req *customizemarket.TaobaoMarketPictureGetuserpicturesAPIRequest, resp *customizemarket.TaobaoMarketPictureGetuserpicturesAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
