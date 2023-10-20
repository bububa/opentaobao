package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamsfserviceauditsupdateAPIResponse 操作改约审批单 API返回值
// alibaba.msfservice.audits.update
//
// 操作改约审批单
type AlibabamsfserviceauditsupdateAPIResponse struct {
	model.CommonResponse
	AlibabamsfserviceauditsupdateAPIResponseModel
}

// AlibabamsfserviceauditsupdateAPIResponseModel is 操作改约审批单 成功返回结果
type AlibabamsfserviceauditsupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_msfservice_audits_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *AlibabamsfserviceauditsupdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
