package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交楼盘信息 APIResponse
alibaba.alihouse.newhome.project.submit

提交楼盘信息
*/
type AlibabaAlihouseNewhomeProjectSubmitAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeProjectSubmitResponse
}

type AlibabaAlihouseNewhomeProjectSubmitResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_submit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回结果体
    
    Result   *AlibabaAlihouseNewhomeProjectSubmitResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
