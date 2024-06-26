package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelIntlRateUpdateAPIResponse 不落库商家推送更新酒店rate API返回值
// taobao.xhotel.intl.rate.update
//
// 商家主动推送不落库商品的酒店信息
type TaobaoXhotelIntlRateUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelIntlRateUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelIntlRateUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelIntlRateUpdateAPIResponseModel).Reset()
}

// TaobaoXhotelIntlRateUpdateAPIResponseModel is 不落库商家推送更新酒店rate 成功返回结果
type TaobaoXhotelIntlRateUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_intl_rate_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoXhotelIntlRateUpdateResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelIntlRateUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelIntlRateUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelIntlRateUpdateAPIResponse)
	},
}

// GetTaobaoXhotelIntlRateUpdateAPIResponse 从 sync.Pool 获取 TaobaoXhotelIntlRateUpdateAPIResponse
func GetTaobaoXhotelIntlRateUpdateAPIResponse() *TaobaoXhotelIntlRateUpdateAPIResponse {
	return poolTaobaoXhotelIntlRateUpdateAPIResponse.Get().(*TaobaoXhotelIntlRateUpdateAPIResponse)
}

// ReleaseTaobaoXhotelIntlRateUpdateAPIResponse 将 TaobaoXhotelIntlRateUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelIntlRateUpdateAPIResponse(v *TaobaoXhotelIntlRateUpdateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelIntlRateUpdateAPIResponse.Put(v)
}
