package drug

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据用户id获取店铺id APIResponse
taobao.alihealth.drug.user.shop.get

提供给千牛智能客服，获取用户当前咨询的店铺ID
*/
type TaobaoAlihealthDrugUserShopGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alihealth_drug_user_shop_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // shopId
    
    ShopId   int64 `json:"shop_id,omitempty" xml:"