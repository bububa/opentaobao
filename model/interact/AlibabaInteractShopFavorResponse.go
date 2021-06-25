package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
店铺收藏 APIResponse
alibaba.interact.shop.favor

店铺收藏mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 酒仙。
*/
type AlibabaInteractShopFavorAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractShopFavorResponse `json:"alibaba_interact_shop_favor_response,omitempty"`
}

type AlibabaInteractShopFavorResponse struct {

    // result=0
    Result   string `json:"result,omitempty"`

}
