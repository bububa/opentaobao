package elife

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoElifeLifecardQueryAPIResponse 查询交易结果 API返回值
// taobao.elife.lifecard.query
//
// 卖家在交易状态不明的情况下, 查询交易结果.
type TaobaoElifeLifecardQueryAPIResponse struct {
	model.CommonResponse
	TaobaoElifeLifecardQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoElifeLifecardQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoElifeLifecardQueryAPIResponseModel).Reset()
}

// TaobaoElifeLifecardQueryAPIResponseModel is 查询交易结果 成功返回结果
type TaobaoElifeLifecardQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"elife_lifecard_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultMsg
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// resultCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// amount
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// inflateAmount
	InflateAmount int64 `json:"inflate_amount,omitempty" xml:"inflate_amount,omitempty"`
	// successed
	Successed bool `json:"successed,omitempty" xml:"successed,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoElifeLifecardQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = ""
	m.Amount = 0
	m.InflateAmount = 0
	m.Successed = false
}

var poolTaobaoElifeLifecardQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoElifeLifecardQueryAPIResponse)
	},
}

// GetTaobaoElifeLifecardQueryAPIResponse 从 sync.Pool 获取 TaobaoElifeLifecardQueryAPIResponse
func GetTaobaoElifeLifecardQueryAPIResponse() *TaobaoElifeLifecardQueryAPIResponse {
	return poolTaobaoElifeLifecardQueryAPIResponse.Get().(*TaobaoElifeLifecardQueryAPIResponse)
}

// ReleaseTaobaoElifeLifecardQueryAPIResponse 将 TaobaoElifeLifecardQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoElifeLifecardQueryAPIResponse(v *TaobaoElifeLifecardQueryAPIResponse) {
	v.Reset()
	poolTaobaoElifeLifecardQueryAPIResponse.Put(v)
}
