package alihealthcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabafmhealthbutlerenergysyncAPIResponse 同步用户消耗能量 API返回值
// alibaba.fmhealth.butler.energy.sync
//
// 同步用户消耗能量，用户消耗s点或卡路里后，同步给健康平台
type AlibabafmhealthbutlerenergysyncAPIResponse struct {
	model.CommonResponse
	AlibabafmhealthbutlerenergysyncAPIResponseModel
}

// AlibabafmhealthbutlerenergysyncAPIResponseModel is 同步用户消耗能量 成功返回结果
type AlibabafmhealthbutlerenergysyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fmhealth_butler_energy_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
