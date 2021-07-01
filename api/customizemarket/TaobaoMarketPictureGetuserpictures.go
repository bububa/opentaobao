package customizemarket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/customizemarket"
)

/* TaobaoMarketPictureGetuserpictures
读取用户上传图片
taobao.market.picture.getuserpictures

商家通过用户信息，获取用户上传的 */
func TaobaoMarketPictureGetuserpictures(clt *core.SDKClient, req *customizemarket.TaobaoMarketPictureGetuserpicturesAPIRequest, session string) (*customizemarket.TaobaoMarketPictureGetuserpicturesAPIResponse, error) {
	var resp customizemarket.TaobaoMarketPictureGetuserpicturesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
