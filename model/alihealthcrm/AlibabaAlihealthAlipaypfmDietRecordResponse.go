package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户每日摄入卡路里总量回传接口 APIResponse
alibaba.alihealth.alipaypfm.diet.record

用户每日摄入卡路里总量回传接口
*/
type AlibabaAlihealthAlipaypfmDietRecordAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthAlipaypfmDietRecordResponse
}

type AlibabaAlihealthAlipaypfmDietRecordResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_alipaypfm_diet_record_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
