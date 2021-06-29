package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交预售证 APIResponse
alibaba.alihouse.newhome.project.presalepermit.submit

提交楼盘预售证信息
*/
type AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeProjectPresalepermitSubmitResponse
}

type AlibabaAlihouseNewhomeProjectPresalepermitSubmitResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_presalepermit_submit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *AlibabaAlihouseNewhomeProjectPresalepermitSubmitResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
