package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse 提交楼盘电话 API返回值
// alibaba.alihouse.newhome.project.phone.submit
//
// 提交楼盘电话
type AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponseModel
}

// AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponseModel is 提交楼盘电话 成功返回结果
type AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_phone_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectPhoneSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
