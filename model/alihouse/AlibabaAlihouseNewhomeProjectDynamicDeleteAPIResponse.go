package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectDynamicDeleteAPIResponse 楼盘快讯删除 API返回值
// alibaba.alihouse.newhome.project.dynamic.delete
//
// 楼盘快讯删除
type AlibabaAlihouseNewhomeProjectDynamicDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectDynamicDeleteAPIResponseModel
}

// AlibabaAlihouseNewhomeProjectDynamicDeleteAPIResponseModel is 楼盘快讯删除 成功返回结果
type AlibabaAlihouseNewhomeProjectDynamicDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_dynamic_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectDynamicDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
