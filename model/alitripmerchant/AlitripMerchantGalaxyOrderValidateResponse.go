package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-订单试单接口 API返回值 
alitrip.merchant.galaxy.order.validate

根据用户选择酒店房型、入住人数、预订时间参数，获取是否可预订及价格变化信息
*/
type AlitripMerchantGalaxyOrderValidateAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyOrderValidateResponse
}

// 星河-订单试单接口 成功返回结果
type AlitripMerchantGalaxyOrderValidateResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_validate_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 默认描述
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`
}
