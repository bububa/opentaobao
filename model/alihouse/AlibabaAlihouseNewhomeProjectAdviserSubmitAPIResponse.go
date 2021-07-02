package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectAdviserSubmitAPIResponse 提交楼盘顾问 API返回值
// alibaba.alihouse.newhome.project.adviser.submit
//
// 提交楼盘顾问
type AlibabaAlihouseNewhomeProjectAdviserSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectAdviserSubmitAPIResponseModel
}

// AlibabaAlihouseNewhomeProjectAdviserSubmitAPIResponseModel is 提交楼盘顾问 成功返回结果
type AlibabaAlihouseNewhomeProjectAdviserSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_adviser_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectAdviserSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
