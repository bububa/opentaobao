package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectkanamequeryAPIResponse 查询KA楼盘名称 API返回值
// alibaba.alihouse.newhome.project.kaname.query
//
// 查询KA楼盘名称
type AlibabaalihousenewhomeprojectkanamequeryAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojectkanamequeryAPIResponseModel
}

// AlibabaalihousenewhomeprojectkanamequeryAPIResponseModel is 查询KA楼盘名称 成功返回结果
type AlibabaalihousenewhomeprojectkanamequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_kaname_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeprojectkanamequeryResult `json:"result,omitempty" xml:"result,omitempty"`
}
