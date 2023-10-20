package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightcitysuggestAPIResponse 机票城市搜索 API返回值
// alitrip.btrip.flight.city.suggest
//
// 提供机票城市搜索接口，提高OA用户对接效率
type AlitripbtripflightcitysuggestAPIResponse struct {
	model.CommonResponse
	AlitripbtripflightcitysuggestAPIResponseModel
}

// AlitripbtripflightcitysuggestAPIResponseModel is 机票城市搜索 成功返回结果
type AlitripbtripflightcitysuggestAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_city_suggest_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BtripApplyResult `json:"result,omitempty" xml:"result,omitempty"`
}
