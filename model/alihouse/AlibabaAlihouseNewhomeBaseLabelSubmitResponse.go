package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交标签库 API返回值 
alibaba.alihouse.newhome.base.label.submit

提交标签库
*/
type AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeBaseLabelSubmitResponse
}

// 提交标签库 成功返回结果
type AlibabaAlihouseNewhomeBaseLabelSubmitResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_base_label_submit_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaAlihouseNewhomeBaseLabelSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
