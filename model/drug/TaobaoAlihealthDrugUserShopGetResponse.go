package drug

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据用户id获取店铺id APIResponse
taobao.alihealth.drug.user.shop.get

提供给千牛智能客服，获取用户当前咨询的店铺ID
*/
type TaobaoAlihealthDrugUserShopGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAlihealthDrugUserShopGetResponse `json:"taobao_alihealth_drug_user_shop_get_response,omitempty"`
}

type TaobaoAlihealthDrugUserShopGetResponse struct {

    // shopId
    ShopId   int64 `json:"shop_id,omitempty"`

}
