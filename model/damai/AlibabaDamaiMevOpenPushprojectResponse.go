package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-项目接口pushProject APIResponse
alibaba.damai.mev.open.pushproject

pushProject
*/
type AlibabaDamaiMevOpenPushprojectAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenPushprojectResponse
}

type AlibabaDamaiMevOpenPushprojectResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_pushproject_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaDamaiMevOpenPushprojectResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
