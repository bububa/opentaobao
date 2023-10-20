package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformservicecapacitysaveAPIResponse 保存服务容量 API返回值
// alibaba.ssc.supplyplatform.servicecapacity.save
//
// 保存服务容量
type AlibabasscsupplyplatformservicecapacitysaveAPIResponse struct {
	model.CommonResponse
	AlibabasscsupplyplatformservicecapacitysaveAPIResponseModel
}

// AlibabasscsupplyplatformservicecapacitysaveAPIResponseModel is 保存服务容量 成功返回结果
type AlibabasscsupplyplatformservicecapacitysaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_servicecapacity_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabasscsupplyplatformservicecapacitysaveResult `json:"result,omitempty" xml:"result,omitempty"`
}
