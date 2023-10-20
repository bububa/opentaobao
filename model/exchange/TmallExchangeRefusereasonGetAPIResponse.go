package exchange

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeRefusereasonGetAPIResponse 获取拒绝换货原因列表 API返回值
// tmall.exchange.refusereason.get
//
// 获取拒绝换货原因列表
type TmallExchangeRefusereasonGetAPIResponse struct {
	model.CommonResponse
	TmallExchangeRefusereasonGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallExchangeRefusereasonGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallExchangeRefusereasonGetAPIResponseModel).Reset()
}

// TmallExchangeRefusereasonGetAPIResponseModel is 获取拒绝换货原因列表 成功返回结果
type TmallExchangeRefusereasonGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_refusereason_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TmallExchangeRefusereasonGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallExchangeRefusereasonGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallExchangeRefusereasonGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallExchangeRefusereasonGetAPIResponse)
	},
}

// GetTmallExchangeRefusereasonGetAPIResponse 从 sync.Pool 获取 TmallExchangeRefusereasonGetAPIResponse
func GetTmallExchangeRefusereasonGetAPIResponse() *TmallExchangeRefusereasonGetAPIResponse {
	return poolTmallExchangeRefusereasonGetAPIResponse.Get().(*TmallExchangeRefusereasonGetAPIResponse)
}

// ReleaseTmallExchangeRefusereasonGetAPIResponse 将 TmallExchangeRefusereasonGetAPIResponse 保存到 sync.Pool
func ReleaseTmallExchangeRefusereasonGetAPIResponse(v *TmallExchangeRefusereasonGetAPIResponse) {
	v.Reset()
	poolTmallExchangeRefusereasonGetAPIResponse.Put(v)
}
