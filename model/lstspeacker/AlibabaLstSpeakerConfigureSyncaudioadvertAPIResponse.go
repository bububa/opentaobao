package lstspeacker

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同步广告 API返回值 
alibaba.lst.speaker.configure.syncaudioadvert

如意音箱广告同步
*/
type AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse struct {
    model.CommonResponse
    AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponseModel
}

// 同步广告 成功返回结果
type AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_lst_speaker_configure_syncaudioadvert_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 执行结果
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`
    // 执行结果标识
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`
    // 错误码
    ErroMessage   string `json:"erro_message,omitempty" xml:"erro_message,omitempty"`
    // 错误码
    ErroCode   string `json:"erro_code,omitempty" xml:"erro_code,omitempty"`
}
