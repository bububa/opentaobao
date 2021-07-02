package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServicestoreSaveAPIResponse 保存网点 API返回值
// alibaba.ssc.supplyplatform.servicestore.save
//
// 网点创建、修改
type AlibabaSscSupplyplatformServicestoreSaveAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServicestoreSaveAPIResponseModel
}

// AlibabaSscSupplyplatformServicestoreSaveAPIResponseModel is 保存网点 成功返回结果
type AlibabaSscSupplyplatformServicestoreSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_servicestore_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaSscSupplyplatformServicestoreSaveResult `json:"result,omitempty" xml:"result,omitempty"`
}
