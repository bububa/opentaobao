package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家同意退货 APIResponse
taobao.rp.returngoods.agree

卖家同意退货，支持淘宝和天猫的订单。
*/
type TaobaoRpReturngoodsAgreeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"rp_returngoods_agree_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"