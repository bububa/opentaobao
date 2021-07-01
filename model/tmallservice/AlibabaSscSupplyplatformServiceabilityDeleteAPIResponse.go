package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSscSupplyplatformServiceabilityDeleteAPIResponse
删除服务能力 API返回值
alibaba.ssc.supplyplatform.serviceability.delete

删除服务能力 */
type AlibabaSscSupplyplatformServiceabilityDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServiceabilityDeleteAPIResponseModel
}

// AlibabaSscSupplyplatformServiceabilityDeleteAPIResponseModel is 删除服务能力 成功返回结果
type AlibabaSscSupplyplatformServiceabilityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_serviceability_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaSscSupplyplatformServiceabilityDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
