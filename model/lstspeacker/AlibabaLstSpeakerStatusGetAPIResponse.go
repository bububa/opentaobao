package lstspeacker

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstspeakerstatusgetAPIResponse 音箱设备在线状态 API返回值
// alibaba.lst.speaker.status.get
//
// 音箱设备在线状态查询
type AlibabalstspeakerstatusgetAPIResponse struct {
	model.CommonResponse
	AlibabalstspeakerstatusgetAPIResponseModel
}

// AlibabalstspeakerstatusgetAPIResponseModel is 音箱设备在线状态 成功返回结果
type AlibabalstspeakerstatusgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_speaker_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabalstspeakerstatusgetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
