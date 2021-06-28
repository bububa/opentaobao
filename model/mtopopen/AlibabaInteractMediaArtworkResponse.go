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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_media_artwork_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

}
