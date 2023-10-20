package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectcooperationsubmitAPIResponse 提交KA合作楼盘 API返回值
// alibaba.alihouse.newhome.project.cooperation.submit
//
// 提交KA合作楼盘
type AlibabaalihousenewhomeprojectcooperationsubmitAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojectcooperationsubmitAPIResponseModel
}

// AlibabaalihousenewhomeprojectcooperationsubmitAPIResponseModel is 提交KA合作楼盘 成功返回结果
type AlibabaalihousenewhomeprojectcooperationsubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_cooperation_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeprojectcooperationsubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
