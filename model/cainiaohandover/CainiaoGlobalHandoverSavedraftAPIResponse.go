package cainiaohandover

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalHandoverSavedraftAPIResponse 创建交接单草稿 API返回值
// cainiao.global.handover.savedraft
//
// 提供给ISV通过该接口创建交接单草稿
type CainiaoGlobalHandoverSavedraftAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalHandoverSavedraftAPIResponseModel
}

// CainiaoGlobalHandoverSavedraftAPIResponseModel is 创建交接单草稿 成功返回结果
type CainiaoGlobalHandoverSavedraftAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_handover_savedraft_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *HsfResult `json:"result,omitempty" xml:"result,omitempty"`
}
