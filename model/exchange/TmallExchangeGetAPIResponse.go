package exchange

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeGetAPIResponse 获取单笔换货详情 API返回值
// tmall.exchange.get
//
// 获取单笔换货详情
type TmallExchangeGetAPIResponse struct {
	model.CommonResponse
	TmallExchangeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallExchangeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallExchangeGetAPIResponseModel).Reset()
}

// TmallExchangeGetAPIResponseModel is 获取单笔换货详情 成功返回结果
type TmallExchangeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ExchangeBaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallExchangeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallExchangeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallExchangeGetAPIResponse)
	},
}

// GetTmallExchangeGetAPIResponse 从 sync.Pool 获取 TmallExchangeGetAPIResponse
func GetTmallExchangeGetAPIResponse() *TmallExchangeGetAPIResponse {
	return poolTmallExchangeGetAPIResponse.Get().(*TmallExchangeGetAPIResponse)
}

// ReleaseTmallExchangeGetAPIResponse 将 TmallExchangeGetAPIResponse 保存到 sync.Pool
func ReleaseTmallExchangeGetAPIResponse(v *TmallExchangeGetAPIResponse) {
	v.Reset()
	poolTmallExchangeGetAPIResponse.Put(v)
}
