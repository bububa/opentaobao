package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse 视频专辑场景接入接口 API返回值
// alibaba.ailabs.aligenie.openvideoalbum.scenepush
//
// 视频专辑场景接入接口
type AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponseModel
}

// AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponseModel is 视频专辑场景接入接口 成功返回结果
type AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_openvideoalbum_scenepush_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 描述
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 状态码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}
