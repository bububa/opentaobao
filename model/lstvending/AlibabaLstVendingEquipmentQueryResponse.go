package lstvending

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
自动售卖机设备信息查询 APIResponse
alibaba.lst.vending.equipment.query

为零售通品牌商提供已租赁的设备信息查询。
*/
type AlibabaLstVendingEquipmentQueryAPIResponse struct {
    model.CommonResponse
    AlibabaLstVendingEquipmentQueryResponse
}

type AlibabaLstVendingEquipmentQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_vending_equipment_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果集
    
    Result   *AlibabaLstVendingEquipmentQueryResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
