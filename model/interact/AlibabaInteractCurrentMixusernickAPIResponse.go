package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractCurrentMixusernickAPIResponse
手淘混淆nick开放接口鉴权专用 API返回值
alibaba.interact.current.mixusernick

手淘混淆nick开放接口鉴权专用，无数据输入输出。 */
type AlibabaInteractCurrentMixusernickAPIResponse struct {
	model.CommonResponse
	AlibabaInteractCurrentMixusernickAPIResponseModel
}

// AlibabaInteractCurrentMixusernickAPIResponseModel is 手淘混淆nick开放接口鉴权专用 成功返回结果
type AlibabaInteractCurrentMixusernickAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_current_mixusernick_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
