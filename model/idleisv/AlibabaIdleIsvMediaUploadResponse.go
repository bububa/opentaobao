package idleisv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼服务商-图片上传 APIResponse
alibaba.idle.isv.media.upload

供外部服务商ISV进行闲鱼商品发布时上传商品所需图片
*/
type AlibabaIdleIsvMediaUploadAPIResponse struct {
    model.CommonResponse
    AlibabaIdleIsvMediaUploadResponse
}

type AlibabaIdleIsvMediaUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_isv_media_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应数据
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
