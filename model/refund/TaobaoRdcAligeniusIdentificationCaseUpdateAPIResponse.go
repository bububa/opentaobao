package refund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdcAligeniusIdentificationCaseUpdateAPIResponse
鉴定工单信息同步 API返回值
taobao.rdc.aligenius.identification.case.update

同步商家鉴定工单信息 */
type TaobaoRdcAligeniusIdentificationCaseUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoRdcAligeniusIdentificationCaseUpdateAPIResponseModel
}

// TaobaoRdcAligeniusIdentificationCaseUpdateAPIResponseModel is 鉴定工单信息同步 成功返回结果
type TaobaoRdcAligeniusIdentificationCaseUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"rdc_aligenius_identification_case_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoRdcAligeniusIdentificationCaseUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
