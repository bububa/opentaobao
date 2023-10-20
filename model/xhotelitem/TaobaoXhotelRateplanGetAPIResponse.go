package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateplanGetAPIResponse 价格计划rateplan查询 API返回值
// taobao.xhotel.rateplan.get
//
// 酒店产品库rateplan查询
type TaobaoXhotelRateplanGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRateplanGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRateplanGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRateplanGetAPIResponseModel).Reset()
}

// TaobaoXhotelRateplanGetAPIResponseModel is 价格计划rateplan查询 成功返回结果
type TaobaoXhotelRateplanGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rateplan_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rateplan
	Rateplan *RatePlan `json:"rateplan,omitempty" xml:"rateplan,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRateplanGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Rateplan = nil
}

var poolTaobaoXhotelRateplanGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRateplanGetAPIResponse)
	},
}

// GetTaobaoXhotelRateplanGetAPIResponse 从 sync.Pool 获取 TaobaoXhotelRateplanGetAPIResponse
func GetTaobaoXhotelRateplanGetAPIResponse() *TaobaoXhotelRateplanGetAPIResponse {
	return poolTaobaoXhotelRateplanGetAPIResponse.Get().(*TaobaoXhotelRateplanGetAPIResponse)
}

// ReleaseTaobaoXhotelRateplanGetAPIResponse 将 TaobaoXhotelRateplanGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRateplanGetAPIResponse(v *TaobaoXhotelRateplanGetAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRateplanGetAPIResponse.Put(v)
}
