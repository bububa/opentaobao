package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectadviserbindAPIResponse 置业顾问批量绑定楼盘 API返回值
// alibaba.alihouse.newhome.project.adviser.bind
//
// 置业顾问批量绑定楼盘
type AlibabaalihousenewhomeprojectadviserbindAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojectadviserbindAPIResponseModel
}

// AlibabaalihousenewhomeprojectadviserbindAPIResponseModel is 置业顾问批量绑定楼盘 成功返回结果
type AlibabaalihousenewhomeprojectadviserbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_adviser_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeprojectadviserbindResult `json:"result,omitempty" xml:"result,omitempty"`
}
