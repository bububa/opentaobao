package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthvaccinorderupdateAPIResponse 回传取号状态 API返回值
// alibaba.health.vaccin.order.update
//
// 回传取号状态
type AlibabahealthvaccinorderupdateAPIResponse struct {
	model.CommonResponse
	AlibabahealthvaccinorderupdateAPIResponseModel
}

// AlibabahealthvaccinorderupdateAPIResponseModel is 回传取号状态 成功返回结果
type AlibabahealthvaccinorderupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_order_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 1
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 1
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 1
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
