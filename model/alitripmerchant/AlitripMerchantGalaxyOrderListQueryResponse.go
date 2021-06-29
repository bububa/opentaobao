package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-订单列表查询 APIResponse
alitrip.merchant.galaxy.order.list.query

为C端用户提供酒店预订订单列表查询服务，包括订单支付状态、订单日期
*/
type AlitripMerchantGalaxyOrderListQueryAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyOrderListQueryResponse
}

type AlitripMerchantGalaxyOrderListQueryResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_list_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 默认描述
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}
