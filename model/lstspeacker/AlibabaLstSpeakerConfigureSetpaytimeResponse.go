package lstspeacker

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
音箱播放配置 APIResponse
alibaba.lst.speaker.configure.setpaytime

音箱播放配置
*/
type AlibabaLstSpeakerConfigureSetpaytimeAPIResponse struct {
    model.CommonResponse
    AlibabaLstSpeakerConfigureSetpaytimeResponse
}

type AlibabaLstSpeakerConfigureSetpaytimeResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_speaker_configure_setpaytime_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 执行结果
    
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`

    
    // 执行结果标识
    
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`

    
    // 错误码
    
    ErroMessage   string `json:"erro_message,omitempty" xml:"erro_message,omitempty"`

    
    // 错误码
    
    ErroCode   string `json:"erro_code,omitempty" xml:"erro_code,omitempty"`

    
}
