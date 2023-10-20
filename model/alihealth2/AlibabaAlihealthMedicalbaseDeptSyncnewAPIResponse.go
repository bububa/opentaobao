package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalbasedeptsyncnewAPIResponse 直连科室上传 API返回值
// alibaba.alihealth.medicalbase.dept.syncnew
//
// 直连科室上传接口
type AlibabaalihealthmedicalbasedeptsyncnewAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthmedicalbasedeptsyncnewAPIResponseModel
}

// AlibabaalihealthmedicalbasedeptsyncnewAPIResponseModel is 直连科室上传 成功返回结果
type AlibabaalihealthmedicalbasedeptsyncnewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_dept_syncnew_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
