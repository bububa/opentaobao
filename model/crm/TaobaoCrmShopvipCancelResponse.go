package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
卖家取消店铺vip的优惠 APIResponse
taobao.crm.shopvip.cancel

此接口用于取消VIP优惠
*/
type TaobaoCrmShopvipCancelAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmShopvipCancelResponse `json:"taobao_crm_shopvip_cancel_response,omitempty"`
}

type TaobaoCrmShopvipCancelResponse struct {

    // 返回操作是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
