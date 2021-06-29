package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单据处理状态查询 APIResponse
alibaba.alihealth.drug.kyt.searchstatus

单据处理状态查询
*/
type AlibabaAlihealthDrugKytSearchstatusAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytSearchstatusResponse
}

type AlibabaAlihealthDrugKytSearchstatusResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_searchstatus_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugKytSearchstatusResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
