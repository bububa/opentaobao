package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除卖家地址库 APIResponse
taobao.logistics.address.remove

用此接口删除卖家地址库
*/
type TaobaoLogisticsAddressRemoveAPIResponse struct {
    model.CommonResponse
    Response *TaobaoLogisticsAddressRemoveResponse `json:"taobao_logistics_address_remove_response,omitempty"`
}

type TaobaoLogisticsAddressRemoveResponse struct {

    // 只返回修改日期modify_date
    AddressResult   *AddressResult `json:"address_result,omitempty"`

}
