package alihealthoutflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-药品同步接口 APIResponse
alibaba.alihealth.outflow.drug.saveorupdate

处方外流-药品同步接口
*/
type AlibabaAlihealthOutflowDrugSaveorupdateAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthOutflowDrugSaveorupdateResponse
}

type AlibabaAlihealthOutflowDrugSaveorupdateResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_outflow_drug_saveorupdate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`

    
}
