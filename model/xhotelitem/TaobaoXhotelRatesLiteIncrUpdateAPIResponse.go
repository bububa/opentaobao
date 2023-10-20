package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRatesLiteIncrUpdateAPIResponse 酒店价格库存轻量级增量接口 API返回值
// taobao.xhotel.rates.lite.incr.update
//
// 多个rate的库存房价开关的增量更新接口
type TaobaoXhotelRatesLiteIncrUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRatesLiteIncrUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRatesLiteIncrUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRatesLiteIncrUpdateAPIResponseModel).Reset()
}

// TaobaoXhotelRatesLiteIncrUpdateAPIResponseModel is 酒店价格库存轻量级增量接口 成功返回结果
type TaobaoXhotelRatesLiteIncrUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rates_lite_incr_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoXhotelRatesLiteIncrUpdateResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRatesLiteIncrUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelRatesLiteIncrUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRatesLiteIncrUpdateAPIResponse)
	},
}

// GetTaobaoXhotelRatesLiteIncrUpdateAPIResponse 从 sync.Pool 获取 TaobaoXhotelRatesLiteIncrUpdateAPIResponse
func GetTaobaoXhotelRatesLiteIncrUpdateAPIResponse() *TaobaoXhotelRatesLiteIncrUpdateAPIResponse {
	return poolTaobaoXhotelRatesLiteIncrUpdateAPIResponse.Get().(*TaobaoXhotelRatesLiteIncrUpdateAPIResponse)
}

// ReleaseTaobaoXhotelRatesLiteIncrUpdateAPIResponse 将 TaobaoXhotelRatesLiteIncrUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRatesLiteIncrUpdateAPIResponse(v *TaobaoXhotelRatesLiteIncrUpdateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRatesLiteIncrUpdateAPIResponse.Put(v)
}
