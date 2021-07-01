package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWirelessBuntingShopShorturlCreateAPIRequest
通过店铺id取得短链 API请求
taobao.wireless.bunting.shop.shorturl.create

通过店铺id取得短链 */
type TaobaoWirelessBuntingShopShorturlCreateAPIRequest struct {
	model.Params
	// 商店id
	_shopId string
}

// New
