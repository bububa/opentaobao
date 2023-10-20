package normalvisa

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelNormalvisaGetdetailAPIResponse 获取单笔订单的详情 API返回值
// taobao.alitrip.travel.normalvisa.getdetail
//
// 获取单笔签证的详细记录
type TaobaoAlitripTravelNormalvisaGetdetailAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelNormalvisaGetdetailAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelNormalvisaGetdetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelNormalvisaGetdetailAPIResponseModel).Reset()
}

// TaobaoAlitripTravelNormalvisaGetdetailAPIResponseModel is 获取单笔订单的详情 成功返回结果
type TaobaoAlitripTravelNormalvisaGetdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_normalvisa_getdetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaoAlitripTravelNormalvisaGetdetailResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelNormalvisaGetdetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripTravelNormalvisaGetdetailAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelNormalvisaGetdetailAPIResponse)
	},
}

// GetTaobaoAlitripTravelNormalvisaGetdetailAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelNormalvisaGetdetailAPIResponse
func GetTaobaoAlitripTravelNormalvisaGetdetailAPIResponse() *TaobaoAlitripTravelNormalvisaGetdetailAPIResponse {
	return poolTaobaoAlitripTravelNormalvisaGetdetailAPIResponse.Get().(*TaobaoAlitripTravelNormalvisaGetdetailAPIResponse)
}

// ReleaseTaobaoAlitripTravelNormalvisaGetdetailAPIResponse 将 TaobaoAlitripTravelNormalvisaGetdetailAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelNormalvisaGetdetailAPIResponse(v *TaobaoAlitripTravelNormalvisaGetdetailAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelNormalvisaGetdetailAPIResponse.Put(v)
}
