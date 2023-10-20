package axindata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinPoiDetailQueryAPIResponse 景点poi详情查询-阿信 API返回值
// taobao.alitrip.travel.axin.poi.detail.query
//
// 景点poi详情查询-阿信
type TaobaoAlitripTravelAxinPoiDetailQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinPoiDetailQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinPoiDetailQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelAxinPoiDetailQueryAPIResponseModel).Reset()
}

// TaobaoAlitripTravelAxinPoiDetailQueryAPIResponseModel is 景点poi详情查询-阿信 成功返回结果
type TaobaoAlitripTravelAxinPoiDetailQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_poi_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripTravelAxinPoiDetailQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelAxinPoiDetailQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelAxinPoiDetailQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinPoiDetailQueryAPIResponse)
	},
}

// GetTaobaoAlitripTravelAxinPoiDetailQueryAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelAxinPoiDetailQueryAPIResponse
func GetTaobaoAlitripTravelAxinPoiDetailQueryAPIResponse() *TaobaoAlitripTravelAxinPoiDetailQueryAPIResponse {
	return poolTaobaoAlitripTravelAxinPoiDetailQueryAPIResponse.Get().(*TaobaoAlitripTravelAxinPoiDetailQueryAPIResponse)
}

// ReleaseTaobaoAlitripTravelAxinPoiDetailQueryAPIResponse 将 TaobaoAlitripTravelAxinPoiDetailQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelAxinPoiDetailQueryAPIResponse(v *TaobaoAlitripTravelAxinPoiDetailQueryAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelAxinPoiDetailQueryAPIResponse.Put(v)
}
