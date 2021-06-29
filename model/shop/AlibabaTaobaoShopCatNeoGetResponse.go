package shop

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
店铺mtop接口鉴权虚拟api APIResponse
alibaba.taobao.shop.cat.neo.get

获取优惠券信息，仅作客户端鉴权虚拟api使用
*/
type AlibabaTaobaoShopCatNeoGetAPIResponse struct {
    model.CommonResponse
    AlibabaTaobaoShopCatNeoGetResponse
}

type AlibabaTaobaoShopCatNeoGetResponse struct {
    XMLName xml.Name `xml:"alibaba_taobao_shop_cat_neo_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 客户端鉴权虚拟api
    
    Unnamed   string `json:"unnamed,omitempty" xml:"unnamed,omitempty"`

    
}
