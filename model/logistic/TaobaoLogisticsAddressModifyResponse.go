package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家地址库修改 API返回值 
taobao.logistics.address.modify

卖家地址库修改
*/
type TaobaoLogisticsAddressModifyAPIResponse struct {
    model.CommonResponse
    TaobaoLogisticsAddressModifyResponse
}

// 卖家地址库修改 成功返回结果
type TaobaoLogisticsAddressModifyResponse struct {
    XMLName xml.Name `xml:"logistics_address_modify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 只返回修改时间modify_date
    AddressResult   *AddressResult `json:"address_result,omitempty" xml:"address_result,omitempty"`
}
