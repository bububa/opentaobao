package travel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelItemBaseModifyAPIResponse 【API3.0】度假线路商品编辑接口 API返回值
// taobao.alitrip.travel.item.base.modify
//
// 旅行度假新商品基本信息修改接口 第三版。提供商家通过TOP API方式修改商品除sku外的基本信息。
type TaobaoAlitripTravelItemBaseModifyAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelItemBaseModifyAPIResponseModel
}

// TaobaoAlitripTravelItemBaseModifyAPIResponseModel is 【API3.0】度假线路商品编辑接口 成功返回结果
type TaobaoAlitripTravelItemBaseModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_item_base_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品修改结果
	TravelItem *TopTravelItem `json:"travel_item,omitempty" xml:"travel_item,omitempty"`
}
