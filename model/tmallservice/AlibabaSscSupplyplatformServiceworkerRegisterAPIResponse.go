package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceworkerRegisterAPIResponse 服务商添加工人 API返回值
// alibaba.ssc.supplyplatform.serviceworker.register
//
// 工人注册
type AlibabaSscSupplyplatformServiceworkerRegisterAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServiceworkerRegisterAPIResponseModel
}

// AlibabaSscSupplyplatformServiceworkerRegisterAPIResponseModel is 服务商添加工人 成功返回结果
type AlibabaSscSupplyplatformServiceworkerRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_serviceworker_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用响应model
	Result *AlibabaSscSupplyplatformServiceworkerRegisterResult `json:"result,omitempty" xml:"result,omitempty"`
}
