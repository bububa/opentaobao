package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
直调单据查询 APIResponse
alibaba.alihealth.drug.kyt.destbill.list

为药企提供直调单据查询功能
*/
type AlibabaAlihealthDrugKytDestbillListAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDestbillListResponse
}

type AlibabaAlihealthDrugKytDestbillListResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_destbill_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回result
    
    Result   *AlibabaAlihealthDrugKytDestbillListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
