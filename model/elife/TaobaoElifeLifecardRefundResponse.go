package elife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌惠卡券冲正退还 APIResponse
taobao.elife.lifecard.refund

淘宝生活汇消费卡虚拟卡，线下冲正退货接口
*/
type TaobaoElifeLifecardRefundAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"elife_lifecard_refund_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回码，成功为空
    
    ResultCode   string `json:"result_code,omitempty" xml:"