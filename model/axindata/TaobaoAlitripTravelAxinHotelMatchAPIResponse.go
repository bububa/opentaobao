package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelmatchAPIResponse 酒店匹配接口-阿信 API返回值
// taobao.alitrip.travel.axin.hotel.match
//
// 酒店匹配接口-阿信
type TaobaoalitriptravelaxinhotelmatchAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelaxinhotelmatchAPIResponseModel
}

// TaobaoalitriptravelaxinhotelmatchAPIResponseModel is 酒店匹配接口-阿信 成功返回结果
type TaobaoalitriptravelaxinhotelmatchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_match_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果result
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
