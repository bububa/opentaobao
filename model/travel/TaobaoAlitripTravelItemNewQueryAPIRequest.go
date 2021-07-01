package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelItemNewQueryAPIRequest
【API3.0】新版度假单个商品查询接口 API请求
taobao.alitrip.travel.item.new.query

新版旅行度假新商品查询接口（单个商品查询） */
type TaobaoAlitripTravelItemNewQueryAPIRequest struct {
	model.Params
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
}

// New
