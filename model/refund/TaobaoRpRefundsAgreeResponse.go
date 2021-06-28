package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同意退款 APIResponse
taobao.rp.refunds.agree

卖家同意退款，支持批量退款，只允许子账号操作。淘宝退款一次最多能退20笔，总金额不超过6000元；天猫退款一次最多能退30笔，总金额不超过10000元。
*/
type TaobaoRpRefundsAgreeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"rp_refunds_agree_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 退款操作结果列表
    
    Results   []RefundMappingResult `json:"results,omitempty" xml:"