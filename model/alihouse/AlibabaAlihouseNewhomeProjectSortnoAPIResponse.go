package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectsortnoAPIResponse 新房排序值同步 API返回值
// alibaba.alihouse.newhome.project.sortno
//
// 新房排序值同步
type AlibabaalihousenewhomeprojectsortnoAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojectsortnoAPIResponseModel
}

// AlibabaalihousenewhomeprojectsortnoAPIResponseModel is 新房排序值同步 成功返回结果
type AlibabaalihousenewhomeprojectsortnoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_sortno_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeprojectsortnoResult `json:"result,omitempty" xml:"result,omitempty"`
}
