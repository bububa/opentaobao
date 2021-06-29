package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交楼盘电话 APIResponse
alibaba.alihouse.newhome.project.phone.submit

提交楼盘电话
*/
type AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeProjectPhoneSubmitResponse
}

type AlibabaAlihouseNewhomeProjectPhoneSubmitResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_phone_submit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeProjectPhoneSubmitResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
