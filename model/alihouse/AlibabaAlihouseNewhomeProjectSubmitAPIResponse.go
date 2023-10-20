package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectsubmitAPIResponse 提交楼盘信息 API返回值
// alibaba.alihouse.newhome.project.submit
//
// 提交楼盘信息
type AlibabaalihousenewhomeprojectsubmitAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojectsubmitAPIResponseModel
}

// AlibabaalihousenewhomeprojectsubmitAPIResponseModel is 提交楼盘信息 成功返回结果
type AlibabaalihousenewhomeprojectsubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果体
	Result *AlibabaalihousenewhomeprojectsubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
