package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojecttradeitemAPIResponse 新二租品同步 API返回值
// alibaba.alihouse.newhome.project.tradeitem
//
// 新二品同步
type AlibabaalihousenewhomeprojecttradeitemAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojecttradeitemAPIResponseModel
}

// AlibabaalihousenewhomeprojecttradeitemAPIResponseModel is 新二租品同步 成功返回结果
type AlibabaalihousenewhomeprojecttradeitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_tradeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeprojecttradeitemResult `json:"result,omitempty" xml:"result,omitempty"`
}
