package lstspeacker

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstspeakerconfiguresyncaudioAPIResponse 音频同步 API返回值
// alibaba.lst.speaker.configure.syncaudio
//
// 音频同步
type AlibabalstspeakerconfiguresyncaudioAPIResponse struct {
	model.CommonResponse
	AlibabalstspeakerconfiguresyncaudioAPIResponseModel
}

// AlibabalstspeakerconfiguresyncaudioAPIResponseModel is 音频同步 成功返回结果
type AlibabalstspeakerconfiguresyncaudioAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_speaker_configure_syncaudio_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ErroMessage string `json:"erro_message,omitempty" xml:"erro_message,omitempty"`
	// 错误码
	ErroCode string `json:"erro_code,omitempty" xml:"erro_code,omitempty"`
	// 执行结果
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
	// 执行结果标识
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}
