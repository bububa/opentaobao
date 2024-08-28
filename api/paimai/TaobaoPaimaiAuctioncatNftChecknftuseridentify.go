package paimai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoPaimaiAuctioncatNftChecknftuseridentify 根据用户数字id和身份证号校验该用户是否已实名认证成功
// taobao.paimai.auctioncat.nft.checknftuseridentify
//
// 根据用户数字id和身份证号校验该用户是否已实名认证成功
func TaobaoPaimaiAuctioncatNftChecknftuseridentify(ctx context.Context, clt *core.SDKClient, req *paimai.TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIRequest, resp *paimai.TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
