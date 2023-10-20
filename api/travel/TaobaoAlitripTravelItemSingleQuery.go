package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelItemSingleQuery 【API3.0】度假单个商品查询接口
// taobao.alitrip.travel.item.single.query
//
// 旅行度假新商品查询接口（单个商品查询） 第三版
func TaobaoAlitripTravelItemSingleQuery(clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemSingleQueryAPIRequest, resp *travel.TaobaoAlitripTravelItemSingleQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
