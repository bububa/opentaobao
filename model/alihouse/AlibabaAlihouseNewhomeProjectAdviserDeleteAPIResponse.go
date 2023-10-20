package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectadviserdeleteAPIResponse 删除楼盘顾问 API返回值
// alibaba.alihouse.newhome.project.adviser.delete
//
// 删除楼盘顾问
type AlibabaalihousenewhomeprojectadviserdeleteAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojectadviserdeleteAPIResponseModel
}

// AlibabaalihousenewhomeprojectadviserdeleteAPIResponseModel is 删除楼盘顾问 成功返回结果
type AlibabaalihousenewhomeprojectadviserdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_adviser_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeprojectadviserdeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
