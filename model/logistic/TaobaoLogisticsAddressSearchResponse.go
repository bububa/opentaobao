package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家地址库 APIResponse
taobao.logistics.address.search

通过此接口查询卖家地址库，
*/
type TaobaoLogisticsAddressSearchAPIResponse struct {
    model.CommonResponse
    TaobaoLogisticsAddressSearchResponse
}

type TaobaoLogisticsAddressSearchResponse struct {
    XMLName xml.Name `xml:"logistics_address_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 一组地址库数据，
    
    Addresses   []AddressResult `json:"addresses,omitempty" xml:"addresses>address_result,omitempty"`
    
    
}
