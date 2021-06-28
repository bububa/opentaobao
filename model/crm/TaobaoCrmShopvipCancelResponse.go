package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家取消店铺vip的优惠 APIResponse
taobao.crm.shopvip.cancel

此接口用于取消VIP优惠
*/
type TaobaoCrmShopvipCancelAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"crm_shopvip_cancel_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"