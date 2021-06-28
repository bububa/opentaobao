package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
卖家地址库修改 APIResponse
taobao.logistics.address.modify

卖家地址库修改
*/
type TaobaoLogisticsAddressModifyAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoLogisticsAddressModifyResponse `json:"logistics_address_modify_response,omitempty"` 
    TaobaoLogisticsAddressModifyResponse
}

/* model for simplify = false
type TaobaoLogisticsAddressModifyResponse struct {

    // 只返回修改时间modify_date
    
    AddressResult  *struct {
        AddressResult  *AddressResult `json:"address_result,omitempty"`
    } `json:"address_result,omitempty"`
    

}
*/

type TaobaoLogisticsAddressModifyResponse struct {

    // 只返回修改时间modify_date
    AddressResult   *AddressResult `json:"address_result,omitempty"`

}
