package hotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelCityGetAPIResponse 酒店城市数据获取接口 API返回值
// taobao.xhotel.city.get
//
// 引流API，对外提供酒店城市数据
type TaobaoXhotelCityGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelCityGetAPIResponseModel
}

// TaobaoXhotelCityGetAPIResponseModel is 酒店城市数据获取接口 成功返回结果
type TaobaoXhotelCityGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_city_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 酒店区域的数据列表
	Divisions []HotelDivision `json:"divisions,omitempty" xml:"divisions>hotel_division,omitempty"`
	// 总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 城市数据的版本。所有城市数据有一个统一的版本，与入参start和count无关。 ISV可通过版本判断城市数据是否有更新。判断方法如下：ISV在第一次拉取数据时请将version保存在本地；以后再调用接口时请比较本地version与接口返回的version。如果本地version小于于接口返回version，则说明城市数据有更新；如果本地version等于接口返回version，则说明城市数据无更新，不需要再继续拉取城市数据。
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}
