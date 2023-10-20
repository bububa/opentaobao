package travel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelitemnewqueryAPIResponse 【API3.0】新版度假单个商品查询接口 API返回值
// taobao.alitrip.travel.item.new.query
//
// 新版旅行度假新商品查询接口（单个商品查询）
type TaobaoalitriptravelitemnewqueryAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelitemnewqueryAPIResponseModel
}

// TaobaoalitriptravelitemnewqueryAPIResponseModel is 【API3.0】新版度假单个商品查询接口 成功返回结果
type TaobaoalitriptravelitemnewqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_item_new_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品查询结果
	TravelItem *FullTravelItem `json:"travel_item,omitempty" xml:"travel_item,omitempty"`
}
