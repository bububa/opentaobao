package lstspeacker

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
如意音箱音频文件长传 API返回值 
alibaba.lst.speaker.file.upload

如意音箱音频文件长传
*/
type AlibabaLstSpeakerFileUploadAPIResponse struct {
    model.CommonResponse
    AlibabaLstSpeakerFileUploadResponse
}

// 如意音箱音频文件长传 成功返回结果
type AlibabaLstSpeakerFileUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_speaker_file_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 异步获取历史数据接口返回结果
    Result   *ResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
