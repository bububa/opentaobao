package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformserviceworkerquitservicestoreAPIResponse 工人退出网点 API返回值
// alibaba.ssc.supplyplatform.serviceworker.quitservicestore
//
// 工人退出网点
type AlibabasscsupplyplatformserviceworkerquitservicestoreAPIResponse struct {
	model.CommonResponse
	AlibabasscsupplyplatformserviceworkerquitservicestoreAPIResponseModel
}

// AlibabasscsupplyplatformserviceworkerquitservicestoreAPIResponseModel is 工人退出网点 成功返回结果
type AlibabasscsupplyplatformserviceworkerquitservicestoreAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_serviceworker_quitservicestore_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabasscsupplyplatformserviceworkerquitservicestoreResult `json:"result,omitempty" xml:"result,omitempty"`
}
