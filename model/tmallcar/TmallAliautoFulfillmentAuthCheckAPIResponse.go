package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoFulfillmentAuthCheckAPIResponse 商家鉴权 API返回值
// tmall.aliauto.fulfillment.auth.check
//
// 商家鉴权
type TmallAliautoFulfillmentAuthCheckAPIResponse struct {
	model.CommonResponse
	TmallAliautoFulfillmentAuthCheckAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAliautoFulfillmentAuthCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAliautoFulfillmentAuthCheckAPIResponseModel).Reset()
}

// TmallAliautoFulfillmentAuthCheckAPIResponseModel is 商家鉴权 成功返回结果
type TmallAliautoFulfillmentAuthCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_fulfillment_auth_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AliAutoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallAliautoFulfillmentAuthCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallAliautoFulfillmentAuthCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAliautoFulfillmentAuthCheckAPIResponse)
	},
}

// GetTmallAliautoFulfillmentAuthCheckAPIResponse 从 sync.Pool 获取 TmallAliautoFulfillmentAuthCheckAPIResponse
func GetTmallAliautoFulfillmentAuthCheckAPIResponse() *TmallAliautoFulfillmentAuthCheckAPIResponse {
	return poolTmallAliautoFulfillmentAuthCheckAPIResponse.Get().(*TmallAliautoFulfillmentAuthCheckAPIResponse)
}

// ReleaseTmallAliautoFulfillmentAuthCheckAPIResponse 将 TmallAliautoFulfillmentAuthCheckAPIResponse 保存到 sync.Pool
func ReleaseTmallAliautoFulfillmentAuthCheckAPIResponse(v *TmallAliautoFulfillmentAuthCheckAPIResponse) {
	v.Reset()
	poolTmallAliautoFulfillmentAuthCheckAPIResponse.Put(v)
}
