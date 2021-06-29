package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除楼盘顾问 APIResponse
alibaba.alihouse.newhome.project.adviser.delete

删除楼盘顾问
*/
type AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeProjectAdviserDeleteResponse
}

type AlibabaAlihouseNewhomeProjectAdviserDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_adviser_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeProjectAdviserDeleteResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
