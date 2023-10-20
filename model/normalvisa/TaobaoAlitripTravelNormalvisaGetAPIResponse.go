package normalvisa

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelNormalvisaGetAPIResponse 获取签证记录 API返回值
// taobao.alitrip.travel.normalvisa.get
//
// 用于获取普通签证的记录信息
type TaobaoAlitripTravelNormalvisaGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelNormalvisaGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelNormalvisaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelNormalvisaGetAPIResponseModel).Reset()
}

// TaobaoAlitripTravelNormalvisaGetAPIResponseModel is 获取签证记录 成功返回结果
type TaobaoAlitripTravelNormalvisaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_normalvisa_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaoAlitripTravelNormalvisaGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelNormalvisaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelNormalvisaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelNormalvisaGetAPIResponse)
	},
}

// GetTaobaoAlitripTravelNormalvisaGetAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelNormalvisaGetAPIResponse
func GetTaobaoAlitripTravelNormalvisaGetAPIResponse() *TaobaoAlitripTravelNormalvisaGetAPIResponse {
	return poolTaobaoAlitripTravelNormalvisaGetAPIResponse.Get().(*TaobaoAlitripTravelNormalvisaGetAPIResponse)
}

// ReleaseTaobaoAlitripTravelNormalvisaGetAPIResponse 将 TaobaoAlitripTravelNormalvisaGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelNormalvisaGetAPIResponse(v *TaobaoAlitripTravelNormalvisaGetAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelNormalvisaGetAPIResponse.Put(v)
}
