package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectphonesubmitAPIResponse 提交楼盘电话 API返回值
// alibaba.alihouse.newhome.project.phone.submit
//
// 提交楼盘电话
type AlibabaalihousenewhomeprojectphonesubmitAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojectphonesubmitAPIResponseModel
}

// AlibabaalihousenewhomeprojectphonesubmitAPIResponseModel is 提交楼盘电话 成功返回结果
type AlibabaalihousenewhomeprojectphonesubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_phone_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeprojectphonesubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
