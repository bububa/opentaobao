package mtopopen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
原图相关鉴权接口 APIResponse
alibaba.interact.media.artwork

拍摄并上传原图相关鉴权接口
*/
type AlibabaInteractMediaArtworkAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractMediaArtworkResponse `json:"alibaba_interact_media_artwork_response,omitempty"`
}

type AlibabaInteractMediaArtworkResponse struct {

}
