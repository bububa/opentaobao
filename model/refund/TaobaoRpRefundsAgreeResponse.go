package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
同意退款 APIResponse
taobao.rp.refunds.agree

卖家同意退款，支持批量退款，只允许子账号操作。淘宝退款一次最多能退20笔，总金额不超过6000元；天猫退款一次最多能退30笔，总金额不超过10000元。
*/
type TaobaoRpRefundsAgreeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoRpRefundsAgreeResponse `json:"taobao_rp_refunds_agree_response,omitempty"`
}

type TaobaoRpRefundsAgreeResponse struct {

    // 退款操作结果列表
    Results   []RefundMappingResult `json:"results,omitempty"`

    // 批量退款操作情况，可选值：OP_SUCC（全部成功），SOME_OP_SUCC（部分成功），OP_FAILURE_UE（全部失败）
    MsgCode   string `json:"msg_code,omitempty"`

    // 操作成功
    Succ   bool `json:"succ,omitempty"`

    // 信息
    Message   string `json:"message,omitempty"`

}
