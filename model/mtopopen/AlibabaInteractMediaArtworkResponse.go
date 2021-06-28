package mtopopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
原图相关鉴权接口 APIResponse
alibaba.interact.media.artwork

拍摄并上传原图相关鉴权接口
*/
type AlibabaInteractMediaArtworkAPIResponse struct {
    model.CommonResponse
    AlibabaInteractMediaArtworkResponse
}

type AlibabaInteractMediaArtworkResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_media_artwork_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

}
