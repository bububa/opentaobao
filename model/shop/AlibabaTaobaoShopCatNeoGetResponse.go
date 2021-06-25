package shop

import (
    "github.com/bububa/opentaobao/model"
)

/* 
店铺mtop接口鉴权虚拟api APIResponse
alibaba.taobao.shop.cat.neo.get

获取优惠券信息，仅作客户端鉴权虚拟api使用
*/
type AlibabaTaobaoShopCatNeoGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaTaobaoShopCatNeoGetResponse `json:"alibaba_taobao_shop_cat_neo_get_response,omitempty"`
}

type AlibabaTaobaoShopCatNeoGetResponse struct {

    // 客户端鉴权虚拟api
    Unnamed   string `json:"unnamed,omitempty"`

}
