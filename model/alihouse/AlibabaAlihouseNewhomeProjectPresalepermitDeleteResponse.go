package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除楼盘预售证 APIResponse
alibaba.alihouse.newhome.project.presalepermit.delete

删除楼盘预售证信息
*/
type AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeProjectPresalepermitDeleteResponse
}

type AlibabaAlihouseNewhomeProjectPresalepermitDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_presalepermit_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeProjectPresalepermitDeleteResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
