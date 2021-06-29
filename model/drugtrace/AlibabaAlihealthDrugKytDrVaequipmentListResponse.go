package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取企业冷链设备信息 APIResponse
alibaba.alihealth.drug.kyt.dr.vaequipment.list

获取企业冷链设备信息
*/
type AlibabaAlihealthDrugKytDrVaequipmentListAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDrVaequipmentListResponse
}

type AlibabaAlihealthDrugKytDrVaequipmentListResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_vaequipment_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugKytDrVaequipmentListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
