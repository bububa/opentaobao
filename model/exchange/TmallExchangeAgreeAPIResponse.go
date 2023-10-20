package exchange

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeAgreeAPIResponse 卖家同意换货申请 API返回值
// tmall.exchange.agree
//
// 卖家同意换货申请
type TmallExchangeAgreeAPIResponse struct {
	model.CommonResponse
	TmallExchangeAgreeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallExchangeAgreeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallExchangeAgreeAPIResponseModel).Reset()
}

// TmallExchangeAgreeAPIResponseModel is 卖家同意换货申请 成功返回结果
type TmallExchangeAgreeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_agree_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ExchangeBaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallExchangeAgreeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallExchangeAgreeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallExchangeAgreeAPIResponse)
	},
}

// GetTmallExchangeAgreeAPIResponse 从 sync.Pool 获取 TmallExchangeAgreeAPIResponse
func GetTmallExchangeAgreeAPIResponse() *TmallExchangeAgreeAPIResponse {
	return poolTmallExchangeAgreeAPIResponse.Get().(*TmallExchangeAgreeAPIResponse)
}

// ReleaseTmallExchangeAgreeAPIResponse 将 TmallExchangeAgreeAPIResponse 保存到 sync.Pool
func ReleaseTmallExchangeAgreeAPIResponse(v *TmallExchangeAgreeAPIResponse) {
	v.Reset()
	poolTmallExchangeAgreeAPIResponse.Put(v)
}
