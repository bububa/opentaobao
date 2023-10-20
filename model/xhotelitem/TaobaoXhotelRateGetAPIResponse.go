package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateGetAPIResponse 酒店产品库rate查询 API返回值
// taobao.xhotel.rate.get
//
// 酒店产品库rate查询
type TaobaoXhotelRateGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRateGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRateGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRateGetAPIResponseModel).Reset()
}

// TaobaoXhotelRateGetAPIResponseModel is 酒店产品库rate查询 成功返回结果
type TaobaoXhotelRateGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rate_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rate
	Rate *Rate `json:"rate,omitempty" xml:"rate,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRateGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Rate = nil
}

var poolTaobaoXhotelRateGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRateGetAPIResponse)
	},
}

// GetTaobaoXhotelRateGetAPIResponse 从 sync.Pool 获取 TaobaoXhotelRateGetAPIResponse
func GetTaobaoXhotelRateGetAPIResponse() *TaobaoXhotelRateGetAPIResponse {
	return poolTaobaoXhotelRateGetAPIResponse.Get().(*TaobaoXhotelRateGetAPIResponse)
}

// ReleaseTaobaoXhotelRateGetAPIResponse 将 TaobaoXhotelRateGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRateGetAPIResponse(v *TaobaoXhotelRateGetAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRateGetAPIResponse.Put(v)
}
