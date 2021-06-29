package shop

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
店铺mtop接口鉴权虚拟api API返回值 
alibaba.taobao.shop.cat.neo.get

获取优惠券信息，仅作客户端鉴权虚拟api使用
*/
type AlibabaTaobaoShopCatNeoGetAPIResponse struct {
    model.CommonResponse
    AlibabaTaobaoShopCatNeoGetResponse
}

// 店铺mtop接口鉴权虚拟api 成功返回结果
type AlibabaTaobaoShopCatNeoGetResponse struct {
    XMLName xml.Name `xml:"alibaba_taobao_shop_cat_neo_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 客户端鉴权虚拟api
    Unnamed   string `json:"unnamed,omitempty" xml:"unnamed,omitempty"`
}
