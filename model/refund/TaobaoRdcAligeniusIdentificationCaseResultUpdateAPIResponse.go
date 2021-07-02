package refund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIResponse 鉴定工单结果同步 API返回值
// taobao.rdc.aligenius.identification.case.result.update
//
// 同步鉴定工单结果信息
type TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIResponseModel
}

// TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIResponseModel is 鉴定工单结果同步 成功返回结果
type TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"rdc_aligenius_identification_case_result_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoRdcAligeniusIdentificationCaseResultUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
