package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinPoiSearchAPIResponse 景点poi搜索-阿信 API返回值
// taobao.alitrip.travel.axin.poi.search
//
// 给阿信提供景点poi搜索
type TaobaoAlitripTravelAxinPoiSearchAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinPoiSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinPoiSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinPoiSearchAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinPoiSearchAPIResponseModel is 景点poi搜索-阿信 成功返回结果
type TaobaoAlitripTravelAxinPoiSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_poi_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripTravelAxinPoiSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinPoiSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinPoiSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinPoiSearchAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinPoiSearchAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinPoiSearchAPIResponse
func GetTaobaoAlitripTravelAxinPoiSearchAPIResponse() *TaobaoAlitripTravelAxinPoiSearchAPIResponse {
	return poolTaobaoAlitripTravelAxinPoiSearchAPIResponse.Get().(*TaobaoAlitripTravelAxinPoiSearchAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinPoiSearchAPIResponse 将 TaobaoAlitripTravelAxinPoiSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinPoiSearchAPIResponse(v *TaobaoAlitripTravelAxinPoiSearchAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinPoiSearchAPIResponse.Put(v)
}
