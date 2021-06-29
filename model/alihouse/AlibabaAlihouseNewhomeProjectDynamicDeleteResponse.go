package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
楼盘快讯删除 APIResponse
alibaba.alihouse.newhome.project.dynamic.delete

楼盘快讯删除
*/
type AlibabaAlihouseNewhomeProjectDynamicDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeProjectDynamicDeleteResponse
}

type AlibabaAlihouseNewhomeProjectDynamicDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_dynamic_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeProjectDynamicDeleteResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
