package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量判定服务是否可达 APIResponse
taobao.logistics.address.reachablebatch.get

批量判定服务是否可达
*/
type TaobaoLogisticsAddressReachablebatchGetAPIResponse struct {
    model.CommonResponse
    TaobaoLogisticsAddressReachablebatchGetResponse
}

type TaobaoLogisticsAddressReachablebatchGetResponse struct {
    XMLName xml.Name `xml:"logistics_address_reachablebatch_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 物流是否可达结果列表
    
    ReachableResults   []AddressReachableTopResult `json:"reachable_results,omitempty" xml:"reachable_results>address_reachable_top_result,omitempty"`
    
    
}
