package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询往来单位 APIResponse
alibaba.alihealth.drug.kyt.yy.listparts

查询往来单位列表
*/
type AlibabaAlihealthDrugKytYyListpartsAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytYyListpartsResponse
}

type AlibabaAlihealthDrugKytYyListpartsResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_yy_listparts_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugKytYyListpartsResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
