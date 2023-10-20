package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelcitygetAPIResponse 城市列表信息查询-阿信 API返回值
// taobao.alitrip.travel.axin.hotel.city.get
//
// 阿信城市列表信息查询
type TaobaoalitriptravelaxinhotelcitygetAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelaxinhotelcitygetAPIResponseModel
}

// TaobaoalitriptravelaxinhotelcitygetAPIResponseModel is 城市列表信息查询-阿信 成功返回结果
type TaobaoalitriptravelaxinhotelcitygetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_city_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoalitriptravelaxinhotelcitygetResult `json:"result,omitempty" xml:"result,omitempty"`
}
