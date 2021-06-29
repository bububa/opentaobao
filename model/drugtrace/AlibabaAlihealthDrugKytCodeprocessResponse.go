package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关联关系处理查询 APIResponse
alibaba.alihealth.drug.kyt.codeprocess

关联关系处理查询
*/
type AlibabaAlihealthDrugKytCodeprocessAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytCodeprocessResponse
}

type AlibabaAlihealthDrugKytCodeprocessResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_codeprocess_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugKytCodeprocessResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
