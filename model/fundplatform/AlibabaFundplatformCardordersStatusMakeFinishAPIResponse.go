package fundplatform

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformCardordersStatusMakeFinishAPIResponse
制卡商通知制卡完成 API返回值
alibaba.fundplatform.cardorders.status.make.finish

当制卡完成后，制卡商需要调用该接口，通知我们制卡已完成。 */
type AlibabaFundplatformCardordersStatusMakeFinishAPIResponse struct {
	model.CommonResponse
	AlibabaFundplatformCardordersStatusMakeFinishAPIResponseModel
}

// AlibabaFundplatformCardordersStatusMakeFinishAPIResponseModel is 制卡商通知制卡完成 成功返回结果
type AlibabaFundplatformCardordersStatusMakeFinishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fundplatform_cardorders_status_make_finish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CardMakingInformResponse `json:"result,omitempty" xml:"result,omitempty"`
}
