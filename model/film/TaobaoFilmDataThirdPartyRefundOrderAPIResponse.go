package film

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmDataThirdPartyRefundOrderAPIResponse 退票接口 API返回值
// taobao.film.data.third.party.refund.order
//
// 淘票票第三方退票接口
type TaobaoFilmDataThirdPartyRefundOrderAPIResponse struct {
	model.CommonResponse
	TaobaoFilmDataThirdPartyRefundOrderAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFilmDataThirdPartyRefundOrderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFilmDataThirdPartyRefundOrderAPIResponseModel).Reset()
}

// TaobaoFilmDataThirdPartyRefundOrderAPIResponseModel is 退票接口 成功返回结果
type TaobaoFilmDataThirdPartyRefundOrderAPIResponseModel struct {
	XMLName xml.Name `xml:"film_data_third_party_refund_order_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultGeneralModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFilmDataThirdPartyRefundOrderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFilmDataThirdPartyRefundOrderAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFilmDataThirdPartyRefundOrderAPIResponse)
	},
}

// GetTaobaoFilmDataThirdPartyRefundOrderAPIResponse 从 sync.Pool 获取 TaobaoFilmDataThirdPartyRefundOrderAPIResponse
func GetTaobaoFilmDataThirdPartyRefundOrderAPIResponse() *TaobaoFilmDataThirdPartyRefundOrderAPIResponse {
	return poolTaobaoFilmDataThirdPartyRefundOrderAPIResponse.Get().(*TaobaoFilmDataThirdPartyRefundOrderAPIResponse)
}

// ReleaseTaobaoFilmDataThirdPartyRefundOrderAPIResponse 将 TaobaoFilmDataThirdPartyRefundOrderAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFilmDataThirdPartyRefundOrderAPIResponse(v *TaobaoFilmDataThirdPartyRefundOrderAPIResponse) {
	v.Reset()
	poolTaobaoFilmDataThirdPartyRefundOrderAPIResponse.Put(v)
}
