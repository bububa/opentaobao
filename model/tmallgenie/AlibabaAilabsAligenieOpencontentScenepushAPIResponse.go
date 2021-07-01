package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieOpencontentScenepushAPIResponse
音频场景接入接口 API返回值
alibaba.ailabs.aligenie.opencontent.scenepush

天猫精灵音频挂靠场景接入 */
type AlibabaAilabsAligenieOpencontentScenepushAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieOpencontentScenepushAPIResponseModel
}

// AlibabaAilabsAligenieOpencontentScenepushAPIResponseModel is 音频场景接入接口 成功返回结果
type AlibabaAilabsAligenieOpencontentScenepushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_opencontent_scenepush_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 返回信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
}
