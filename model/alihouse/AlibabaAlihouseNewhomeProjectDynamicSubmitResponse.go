package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交楼盘快讯 API返回值 
alibaba.alihouse.newhome.project.dynamic.submit

提交楼盘快讯
*/
type AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeProjectDynamicSubmitResponse
}

// 提交楼盘快讯 成功返回结果
type AlibabaAlihouseNewhomeProjectDynamicSubmitResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_dynamic_submit_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaAlihouseNewhomeProjectDynamicSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
