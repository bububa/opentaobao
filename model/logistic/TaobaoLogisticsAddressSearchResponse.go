package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家地址库 APIResponse
taobao.logistics.address.search

通过此接口查询卖家地址库，
*/
type TaobaoLogisticsAddressSearchAPIResponse struct {
    model.CommonResponse
    Response *TaobaoLogisticsAddressSearchResponse `json:"taobao_logistics_address_search_response,omitempty"`
}

type TaobaoLogisticsAddressSearchResponse struct {

    // 一组地址库数据，
    Addresses   []AddressResult `json:"addresses,omitempty"`

}
