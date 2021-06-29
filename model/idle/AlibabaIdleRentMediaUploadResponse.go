package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼多媒体上传接口 APIResponse
alibaba.idle.rent.media.upload

上传多媒体信息，包括图片、视频（暂不支持）
*/
type AlibabaIdleRentMediaUploadAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRentMediaUploadResponse
}

type AlibabaIdleRentMediaUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_rent_media_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应数据
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
