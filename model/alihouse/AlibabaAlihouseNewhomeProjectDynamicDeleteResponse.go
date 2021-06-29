package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
楼盘快讯删除 API返回值 
alibaba.alihouse.newhome.project.dynamic.delete

楼盘快讯删除
*/
type AlibabaAlihouseNewhomeProjectDynamicDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeProjectDynamicDeleteResponse
}

// 楼盘快讯删除 成功返回结果
type AlibabaAlihouseNewhomeProjectDynamicDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_dynamic_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaAlihouseNewhomeProjectDynamicDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}