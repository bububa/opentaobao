package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoRefundRefundactionsJudgementAPIResponse 判断当前用户是否能对订单执行一些逆向操作，比如退货操作 API返回值
// cainiao.refund.refundactions.judgement
//
// 判断当前用户是否能对订单执行一些逆向操作，比如退货操作
type CainiaoRefundRefundactionsJudgementAPIResponse struct {
	model.CommonResponse
	CainiaoRefundRefundactionsJudgementAPIResponseModel
}

// CainiaoRefundRefundactionsJudgementAPIResponseModel is 判断当前用户是否能对订单执行一些逆向操作，比如退货操作 成功返回结果
type CainiaoRefundRefundactionsJudgementAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_refund_refundactions_judgement_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *CainiaoRefundRefundactionsJudgementBizResult `json:"result,omitempty" xml:"result,omitempty"`
}
