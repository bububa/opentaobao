package fundplatform

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabafundplatformcardordersstatussendedAPIResponse 制卡商通知实体卡发货完成 API返回值
// alibaba.fundplatform.cardorders.status.sended
//
// 当制卡商将实体卡发货完成后，需要调用该接口，通知我们已发货。
type AlibabafundplatformcardordersstatussendedAPIResponse struct {
	model.CommonResponse
	AlibabafundplatformcardordersstatussendedAPIResponseModel
}

// AlibabafundplatformcardordersstatussendedAPIResponseModel is 制卡商通知实体卡发货完成 成功返回结果
type AlibabafundplatformcardordersstatussendedAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fundplatform_cardorders_status_sended_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CardMakingInformResponse `json:"result,omitempty" xml:"result,omitempty"`
}
