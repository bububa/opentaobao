package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelItemSingleQueryAPIRequest
【API3.0】度假单个商品查询接口 API请求
taobao.alitrip.travel.item.single.query

旅行度假新商品查询接口（单个商品查询） 第三版 */
type TaobaoAlitripTravelItemSingleQueryAPIRequest struct {
	model.Params
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
}

// New
