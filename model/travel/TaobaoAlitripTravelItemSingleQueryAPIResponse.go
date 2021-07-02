package travel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelItemSingleQueryAPIResponse 【API3.0】度假单个商品查询接口 API返回值
// taobao.alitrip.travel.item.single.query
//
// 旅行度假新商品查询接口（单个商品查询） 第三版
type TaobaoAlitripTravelItemSingleQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelItemSingleQueryAPIResponseModel
}

// TaobaoAlitripTravelItemSingleQueryAPIResponseModel is 【API3.0】度假单个商品查询接口 成功返回结果
type TaobaoAlitripTravelItemSingleQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_item_single_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品查询结果
	TravelItem *PontusTravelFullTravelItem `json:"travel_item,omitempty" xml:"travel_item,omitempty"`
}
