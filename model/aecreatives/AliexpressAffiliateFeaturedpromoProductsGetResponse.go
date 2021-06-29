package aecreatives

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟主题推广活动商品信息获取 API返回值 
aliexpress.affiliate.featuredpromo.products.get

根据联盟主题推广活动或主题品库查询对应的商品。如下品库为固定品库，可长期调用。品库类型和名称如下：高佣品（Hot Product）、新品（New Arrival）、热销商品（Best Seller）、每周尖货（weeklydeals）
*/
type AliexpressAffiliateFeaturedpromoProductsGetAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateFeaturedpromoProductsGetResponse
}

// 联盟主题推广活动商品信息获取 成功返回结果
type AliexpressAffiliateFeaturedpromoProductsGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_featuredpromo_products_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    RespResult   *ResponseDTO `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}
