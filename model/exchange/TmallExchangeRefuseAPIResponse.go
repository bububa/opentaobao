package exchange

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeRefuseAPIResponse 卖家拒绝换货申请 API返回值
// tmall.exchange.refuse
//
// 卖家拒绝换货申请
type TmallExchangeRefuseAPIResponse struct {
	model.CommonResponse
	TmallExchangeRefuseAPIResponseModel
}

// Reset 清空结构体
func (m *TmallExchangeRefuseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallExchangeRefuseAPIResponseModel).Reset()
}

// TmallExchangeRefuseAPIResponseModel is 卖家拒绝换货申请 成功返回结果
type TmallExchangeRefuseAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_refuse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ExchangeBaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallExchangeRefuseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallExchangeRefuseAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallExchangeRefuseAPIResponse)
	},
}

// GetTmallExchangeRefuseAPIResponse 从 sync.Pool 获取 TmallExchangeRefuseAPIResponse
func GetTmallExchangeRefuseAPIResponse() *TmallExchangeRefuseAPIResponse {
	return poolTmallExchangeRefuseAPIResponse.Get().(*TmallExchangeRefuseAPIResponse)
}

// ReleaseTmallExchangeRefuseAPIResponse 将 TmallExchangeRefuseAPIResponse 保存到 sync.Pool
func ReleaseTmallExchangeRefuseAPIResponse(v *TmallExchangeRefuseAPIResponse) {
	v.Reset()
	poolTmallExchangeRefuseAPIResponse.Put(v)
}
