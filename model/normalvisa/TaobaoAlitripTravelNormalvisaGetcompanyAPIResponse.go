package normalvisa

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelNormalvisaGetcompanyAPIResponse 获取物流公司信息 API返回值
// taobao.alitrip.travel.normalvisa.getcompany
//
// 获取物流公司信息
type TaobaoAlitripTravelNormalvisaGetcompanyAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelNormalvisaGetcompanyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelNormalvisaGetcompanyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelNormalvisaGetcompanyAPIResponseModel).Reset()
}

// TaobaoAlitripTravelNormalvisaGetcompanyAPIResponseModel is 获取物流公司信息 成功返回结果
type TaobaoAlitripTravelNormalvisaGetcompanyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_normalvisa_getcompany_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果，有返回代表成功
	Result *TaobaoAlitripTravelNormalvisaGetcompanyResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelNormalvisaGetcompanyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelNormalvisaGetcompanyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelNormalvisaGetcompanyAPIResponse)
	},
}

// GetTaobaoAlitripTravelNormalvisaGetcompanyAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelNormalvisaGetcompanyAPIResponse
func GetTaobaoAlitripTravelNormalvisaGetcompanyAPIResponse() *TaobaoAlitripTravelNormalvisaGetcompanyAPIResponse {
	return poolTaobaoAlitripTravelNormalvisaGetcompanyAPIResponse.Get().(*TaobaoAlitripTravelNormalvisaGetcompanyAPIResponse)
}

// ReleaseTaobaoAlitripTravelNormalvisaGetcompanyAPIResponse 将 TaobaoAlitripTravelNormalvisaGetcompanyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelNormalvisaGetcompanyAPIResponse(v *TaobaoAlitripTravelNormalvisaGetcompanyAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelNormalvisaGetcompanyAPIResponse.Put(v)
}
