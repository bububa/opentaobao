package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// Taobaopaimaiauctioncatnftchecknftuseridentify 根据用户数字id和身份证号校验该用户是否已实名认证成功
// taobao.paimai.auctioncat.nft.checknftuseridentify
//
// 根据用户数字id和身份证号校验该用户是否已实名认证成功
func Taobaopaimaiauctioncatnftchecknftuseridentify(clt *core.SDKClient, req *paimai.TaobaopaimaiauctioncatnftchecknftuseridentifyAPIRequest, session string) (*paimai.TaobaopaimaiauctioncatnftchecknftuseridentifyAPIResponse, error) {
	var resp paimai.TaobaopaimaiauctioncatnftchecknftuseridentifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
