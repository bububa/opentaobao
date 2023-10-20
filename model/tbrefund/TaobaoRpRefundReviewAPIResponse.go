package tbrefund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRpRefundReviewAPIResponse 审核退款单 API返回值
// taobao.rp.refund.review
//
// 审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。
type TaobaoRpRefundReviewAPIResponse struct {
	model.CommonResponse
	TaobaoRpRefundReviewAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRpRefundReviewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRpRefundReviewAPIResponseModel).Reset()
}

// TaobaoRpRefundReviewAPIResponseModel is 审核退款单 成功返回结果
type TaobaoRpRefundReviewAPIResponseModel struct {
	XMLName xml.Name `xml:"rp_refund_review_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRpRefundReviewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoRpRefundReviewAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRpRefundReviewAPIResponse)
	},
}

// GetTaobaoRpRefundReviewAPIResponse 从 sync.Pool 获取 TaobaoRpRefundReviewAPIResponse
func GetTaobaoRpRefundReviewAPIResponse() *TaobaoRpRefundReviewAPIResponse {
	return poolTaobaoRpRefundReviewAPIResponse.Get().(*TaobaoRpRefundReviewAPIResponse)
}

// ReleaseTaobaoRpRefundReviewAPIResponse 将 TaobaoRpRefundReviewAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRpRefundReviewAPIResponse(v *TaobaoRpRefundReviewAPIResponse) {
	v.Reset()
	poolTaobaoRpRefundReviewAPIResponse.Put(v)
}
