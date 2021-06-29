package damai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-项目接口deleteProject APIResponse
alibaba.damai.mev.open.deleteproject

deleteProject
*/
type AlibabaDamaiMevOpenDeleteprojectAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMevOpenDeleteprojectResponse
}

type AlibabaDamaiMevOpenDeleteprojectResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_mev_open_deleteproject_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaDamaiMevOpenDeleteprojectResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
