package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformserviceworkeravailableworkerAPIResponse 查询可用工人 API返回值
// alibaba.ssc.supplyplatform.serviceworker.availableworker
//
// 可用工人查询
type AlibabasscsupplyplatformserviceworkeravailableworkerAPIResponse struct {
	model.CommonResponse
	AlibabasscsupplyplatformserviceworkeravailableworkerAPIResponseModel
}

// AlibabasscsupplyplatformserviceworkeravailableworkerAPIResponseModel is 查询可用工人 成功返回结果
type AlibabasscsupplyplatformserviceworkeravailableworkerAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_serviceworker_availableworker_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 统一响应对象模型
	Result *AlibabasscsupplyplatformserviceworkeravailableworkerResult `json:"result,omitempty" xml:"result,omitempty"`
}
