package idleitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
图片上传 APIResponse
alibaba.idle.item.media.add

上传图片
*/
type AlibabaIdleItemMediaAddAPIResponse struct {
    model.CommonResponse
    AlibabaIdleItemMediaAddResponse
}

type AlibabaIdleItemMediaAddResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_item_media_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *EasyResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
