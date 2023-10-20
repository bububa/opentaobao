package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectsalestimeAPIResponse 楼盘销售时刻同步 API返回值
// alibaba.alihouse.newhome.project.salestime
//
// 楼盘销售时刻同步
type AlibabaalihousenewhomeprojectsalestimeAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojectsalestimeAPIResponseModel
}

// AlibabaalihousenewhomeprojectsalestimeAPIResponseModel is 楼盘销售时刻同步 成功返回结果
type AlibabaalihousenewhomeprojectsalestimeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_salestime_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeprojectsalestimeResult `json:"result,omitempty" xml:"result,omitempty"`
}
