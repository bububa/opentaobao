package elife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌惠卡券核销 APIResponse
taobao.elife.lifecard.consume

用户线上购买生活汇品牌惠虚拟消费卡，线下购物时，商家码枪核销，涉及用户虚拟卡余额扣减操作
*/
type TaobaoElifeLifecardConsumeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"elife_lifecard_consume_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 本金
    
    Amount   int64 `json:"amount,omitempty" xml:"