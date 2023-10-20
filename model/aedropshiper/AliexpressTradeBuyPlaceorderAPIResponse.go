package aedropshiper

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressTradeBuyPlaceorderAPIResponse AE下单API API返回值
// aliexpress.trade.buy.placeorder
//
// 150欧欧盟税改
type AliexpressTradeBuyPlaceorderAPIResponse struct {
	model.CommonResponse
	AliexpressTradeBuyPlaceorderAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressTradeBuyPlaceorderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressTradeBuyPlaceorderAPIResponseModel).Reset()
}

// AliexpressTradeBuyPlaceorderAPIResponseModel is AE下单API 成功返回结果
type AliexpressTradeBuyPlaceorderAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_trade_buy_placeorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PlaceOrderRes4OpenApiDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressTradeBuyPlaceorderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressTradeBuyPlaceorderAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressTradeBuyPlaceorderAPIResponse)
	},
}

// GetAliexpressTradeBuyPlaceorderAPIResponse 从 sync.Pool 获取 AliexpressTradeBuyPlaceorderAPIResponse
func GetAliexpressTradeBuyPlaceorderAPIResponse() *AliexpressTradeBuyPlaceorderAPIResponse {
	return poolAliexpressTradeBuyPlaceorderAPIResponse.Get().(*AliexpressTradeBuyPlaceorderAPIResponse)
}

// ReleaseAliexpressTradeBuyPlaceorderAPIResponse 将 AliexpressTradeBuyPlaceorderAPIResponse 保存到 sync.Pool
func ReleaseAliexpressTradeBuyPlaceorderAPIResponse(v *AliexpressTradeBuyPlaceorderAPIResponse) {
	v.Reset()
	poolAliexpressTradeBuyPlaceorderAPIResponse.Put(v)
}
