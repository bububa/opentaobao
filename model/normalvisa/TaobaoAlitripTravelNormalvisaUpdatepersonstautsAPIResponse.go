package normalvisa

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse 更新签证办理进度 API返回值
// taobao.alitrip.travel.normalvisa.updatepersonstauts
//
// 更新签证办理进度
type TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponseModel).Reset()
}

// TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponseModel is 更新签证办理进度 成功返回结果
type TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_normalvisa_updatepersonstauts_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果：包含result更新成功
	Result *TaobaoAlitripTravelNormalvisaUpdatepersonstautsResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse)
	},
}

// GetTaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse
func GetTaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse() *TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse {
	return poolTaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse.Get().(*TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse)
}

// ReleaseTaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse 将 TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse(v *TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse.Put(v)
}
