package customizemarket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMarketPictureGetuserpicturesAPIRequest
读取用户上传图片 API请求
taobao.market.picture.getuserpictures

商家通过用户信息，获取用户上传的 */
type TaobaoMarketPictureGetuserpicturesAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
}

// New
