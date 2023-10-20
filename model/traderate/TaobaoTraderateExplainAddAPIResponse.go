package traderate

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTraderateExplainAddAPIResponse 商城评价解释接口 API返回值
// taobao.traderate.explain.add
//
// 商城卖家给评价做出解释（买家追加评论后、评价超过30天的，都不能再做评价解释）
type TaobaoTraderateExplainAddAPIResponse struct {
	model.CommonResponse
	TaobaoTraderateExplainAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTraderateExplainAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTraderateExplainAddAPIResponseModel).Reset()
}

// TaobaoTraderateExplainAddAPIResponseModel is 商城评价解释接口 成功返回结果
type TaobaoTraderateExplainAddAPIResponseModel struct {
	XMLName xml.Name `xml:"traderate_explain_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商城卖家给评价解释是否成功。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTraderateExplainAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoTraderateExplainAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTraderateExplainAddAPIResponse)
	},
}

// GetTaobaoTraderateExplainAddAPIResponse 从 sync.Pool 获取 TaobaoTraderateExplainAddAPIResponse
func GetTaobaoTraderateExplainAddAPIResponse() *TaobaoTraderateExplainAddAPIResponse {
	return poolTaobaoTraderateExplainAddAPIResponse.Get().(*TaobaoTraderateExplainAddAPIResponse)
}

// ReleaseTaobaoTraderateExplainAddAPIResponse 将 TaobaoTraderateExplainAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTraderateExplainAddAPIResponse(v *TaobaoTraderateExplainAddAPIResponse) {
	v.Reset()
	poolTaobaoTraderateExplainAddAPIResponse.Put(v)
}
