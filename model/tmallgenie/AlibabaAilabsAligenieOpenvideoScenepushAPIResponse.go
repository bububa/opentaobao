package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieOpenvideoScenepushAPIResponse
视频单集场景接入API API返回值
alibaba.ailabs.aligenie.openvideo.scenepush

视频单集场景接入API */
type AlibabaAilabsAligenieOpenvideoScenepushAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieOpenvideoScenepushAPIResponseModel
}

// AlibabaAilabsAligenieOpenvideoScenepushAPIResponseModel is 视频单集场景接入API 成功返回结果
type AlibabaAilabsAligenieOpenvideoScenepushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_openvideo_scenepush_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 描述
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
}
