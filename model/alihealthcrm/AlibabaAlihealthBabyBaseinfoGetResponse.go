package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
三方从我们这获取宝宝信息 APIResponse
alibaba.alihealth.baby.baseinfo.get

三方从我们这获取宝宝信息
*/
type AlibabaAlihealthBabyBaseinfoGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthBabyBaseinfoGetResponse
}

type AlibabaAlihealthBabyBaseinfoGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_baby_baseinfo_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BabyInfoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
