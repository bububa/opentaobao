package tbrefund

import (
	"encoding/xml"

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

// TaobaoRpRefundReviewAPIResponseModel is 审核退款单 成功返回结果
type TaobaoRpRefundReviewAPIResponseModel struct {
	XMLName xml.Name `xml:"rp_refund_review_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
