package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交楼盘顾问 APIResponse
alibaba.alihouse.newhome.project.adviser.submit

提交楼盘顾问
*/
type AlibabaAlihouseNewhomeProjectAdviserSubmitAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeProjectAdviserSubmitResponse
}

type AlibabaAlihouseNewhomeProjectAdviserSubmitResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_adviser_submit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeProjectAdviserSubmitResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
