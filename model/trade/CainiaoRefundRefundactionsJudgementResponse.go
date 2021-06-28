package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
判断当前用户是否能对订单执行一些逆向操作，比如退货操作 APIResponse
cainiao.refund.refundactions.judgement

判断当前用户是否能对订单执行一些逆向操作，比如退货操作
*/
type CainiaoRefundRefundactionsJudgementAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_refund_refundactions_judgement_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果对象
    
    Result   *CainiaoRefundRefundactionsJudgementBizResult `json:"result,omitempty" xml:"