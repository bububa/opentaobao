package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交标签库 APIResponse
alibaba.alihouse.newhome.base.label.submit

提交标签库
*/
type AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeBaseLabelSubmitResponse
}

type AlibabaAlihouseNewhomeBaseLabelSubmitResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_base_label_submit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeBaseLabelSubmitResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
