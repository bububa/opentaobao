package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家回填物流信息 APIResponse
taobao.rp.returngoods.refill

卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。
*/
type TaobaoRpReturngoodsRefillAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"rp_returngoods_refill_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 验货操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"