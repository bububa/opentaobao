package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectadvisersubmitAPIResponse 提交楼盘顾问 API返回值
// alibaba.alihouse.newhome.project.adviser.submit
//
// 提交楼盘顾问
type AlibabaalihousenewhomeprojectadvisersubmitAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojectadvisersubmitAPIResponseModel
}

// AlibabaalihousenewhomeprojectadvisersubmitAPIResponseModel is 提交楼盘顾问 成功返回结果
type AlibabaalihousenewhomeprojectadvisersubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_adviser_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeprojectadvisersubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
