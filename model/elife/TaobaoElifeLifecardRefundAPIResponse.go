package elife

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoElifeLifecardRefundAPIResponse 品牌惠卡券冲正退还 API返回值
// taobao.elife.lifecard.refund
//
// 淘宝生活汇消费卡虚拟卡，线下冲正退货接口
type TaobaoElifeLifecardRefundAPIResponse struct {
	model.CommonResponse
	TaobaoElifeLifecardRefundAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoElifeLifecardRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoElifeLifecardRefundAPIResponseModel).Reset()
}

// TaobaoElifeLifecardRefundAPIResponseModel is 品牌惠卡券冲正退还 成功返回结果
type TaobaoElifeLifecardRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"elife_lifecard_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回码，成功为空
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 返回信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 本金
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 膨胀金
	InflateAmount int64 `json:"inflate_amount,omitempty" xml:"inflate_amount,omitempty"`
	// 成功失败标志
	Successed bool `json:"successed,omitempty" xml:"successed,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoElifeLifecardRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.Amount = 0
	m.InflateAmount = 0
	m.Successed = false
}

var poolTaobaoElifeLifecardRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoElifeLifecardRefundAPIResponse)
	},
}

// GetTaobaoElifeLifecardRefundAPIResponse 从 sync.Pool 获取 TaobaoElifeLifecardRefundAPIResponse
func GetTaobaoElifeLifecardRefundAPIResponse() *TaobaoElifeLifecardRefundAPIResponse {
	return poolTaobaoElifeLifecardRefundAPIResponse.Get().(*TaobaoElifeLifecardRefundAPIResponse)
}

// ReleaseTaobaoElifeLifecardRefundAPIResponse 将 TaobaoElifeLifecardRefundAPIResponse 保存到 sync.Pool
func ReleaseTaobaoElifeLifecardRefundAPIResponse(v *TaobaoElifeLifecardRefundAPIResponse) {
	v.Reset()
	poolTaobaoElifeLifecardRefundAPIResponse.Put(v)
}
