package travel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelGereralskuUpdateAPIResponse 发布SKU信息（如果properties重复 则更新） API返回值
// alitrip.travel.gereralsku.update
//
// 发布SKU信息（如果properties重复 则更新）
type AlitripTravelGereralskuUpdateAPIResponse struct {
	model.CommonResponse
	AlitripTravelGereralskuUpdateAPIResponseModel
}

// AlitripTravelGereralskuUpdateAPIResponseModel is 发布SKU信息（如果properties重复 则更新） 成功返回结果
type AlitripTravelGereralskuUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_gereralsku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	FirstResult *TopTravelItem `json:"first_result,omitempty" xml:"first_result,omitempty"`
}
