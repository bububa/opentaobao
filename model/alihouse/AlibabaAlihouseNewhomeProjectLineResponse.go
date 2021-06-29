package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
楼盘上下架 APIResponse
alibaba.alihouse.newhome.project.line

上下架楼盘
*/
type AlibabaAlihouseNewhomeProjectLineAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeProjectLineResponse
}

type AlibabaAlihouseNewhomeProjectLineResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_line_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeProjectLineResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
