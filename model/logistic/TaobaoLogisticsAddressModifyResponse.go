package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家地址库修改 APIResponse
taobao.logistics.address.modify

卖家地址库修改
*/
type TaobaoLogisticsAddressModifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"logistics_address_modify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 只返回修改时间modify_date
    
    AddressResult   *AddressResult `json:"address_result,omitempty" xml:"