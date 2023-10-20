package customizemarket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/customizemarket"
)

// Taobaomarketpicturegetuserpictures 读取用户上传图片
// taobao.market.picture.getuserpictures
//
// 商家通过用户信息，获取用户上传的
func Taobaomarketpicturegetuserpictures(clt *core.SDKClient, req *customizemarket.TaobaomarketpicturegetuserpicturesAPIRequest, session string) (*customizemarket.TaobaomarketpicturegetuserpicturesAPIResponse, error) {
	var resp customizemarket.TaobaomarketpicturegetuserpicturesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
