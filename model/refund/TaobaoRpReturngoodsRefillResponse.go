package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
卖家回填物流信息 APIResponse
taobao.rp.returngoods.refill

卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。
*/
type TaobaoRpReturngoodsRefillAPIResponse struct {
    model.CommonResponse
    Response *TaobaoRpReturngoodsRefillResponse `json:"taobao_rp_returngoods_refill_response,omitempty"`
}

type TaobaoRpReturngoodsRefillResponse struct {

    // 验货操作是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
