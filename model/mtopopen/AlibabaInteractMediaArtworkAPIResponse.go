package mtopopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractMediaArtworkAPIResponse
原图相关鉴权接口 API返回值
alibaba.interact.media.artwork

拍摄并上传原图相关鉴权接口 */
type AlibabaInteractMediaArtworkAPIResponse struct {
	model.CommonResponse
	AlibabaInteractMediaArtworkAPIResponseModel
}

// AlibabaInteractMediaArtworkAPIResponseModel is 原图相关鉴权接口 成功返回结果
type AlibabaInteractMediaArtworkAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_media_artwork_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
