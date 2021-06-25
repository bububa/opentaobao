package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量判定服务是否可达 APIResponse
taobao.logistics.address.reachablebatch.get

批量判定服务是否可达
*/
type TaobaoLogisticsAddressReachablebatchGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoLogisticsAddressReachablebatchGetResponse `json:"taobao_logistics_address_reachablebatch_get_response,omitempty"`
}

type TaobaoLogisticsAddressReachablebatchGetResponse struct {

    // 物流是否可达结果列表
    ReachableResults   []AddressReachableTopResult `json:"reachable_results,omitempty"`

}
