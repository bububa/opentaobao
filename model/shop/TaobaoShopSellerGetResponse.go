package shop

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家店铺基础信息查询 API返回值 
taobao.shop.seller.get

获取卖家店铺的基本信息
*/
type TaobaoShopSellerGetAPIResponse struct {
    model.CommonResponse
    TaobaoShopSellerGetResponse
}

// 卖家店铺基础信息查询 成功返回结果
type TaobaoShopSellerGetResponse struct {
    XMLName xml.Name `xml:"shop_seller_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 店铺信息
    Shop   *Shop `json:"shop,omitempty" xml:"shop,omitempty"`
}
