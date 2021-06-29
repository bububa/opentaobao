package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交KA合作楼盘 APIResponse
alibaba.alihouse.newhome.project.cooperation.submit

提交KA合作楼盘
*/
type AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeProjectCooperationSubmitResponse
}

type AlibabaAlihouseNewhomeProjectCooperationSubmitResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_cooperation_submit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeProjectCooperationSubmitResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
