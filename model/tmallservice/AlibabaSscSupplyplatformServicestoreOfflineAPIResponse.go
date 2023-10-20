package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformservicestoreofflineAPIResponse 网点下线 API返回值
// alibaba.ssc.supplyplatform.servicestore.offline
//
// 网点下线功能
type AlibabasscsupplyplatformservicestoreofflineAPIResponse struct {
	model.CommonResponse
	AlibabasscsupplyplatformservicestoreofflineAPIResponseModel
}

// AlibabasscsupplyplatformservicestoreofflineAPIResponseModel is 网点下线 成功返回结果
type AlibabasscsupplyplatformservicestoreofflineAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_servicestore_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabasscsupplyplatformservicestoreofflineResult `json:"result,omitempty" xml:"result,omitempty"`
}
