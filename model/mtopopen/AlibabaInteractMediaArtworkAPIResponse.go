package mtopopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractMediaArtworkAPIResponse 原图相关鉴权接口 API返回值
// alibaba.interact.media.artwork
//
// 拍摄并上传原图相关鉴权接口
type AlibabaInteractMediaArtworkAPIResponse struct {
	model.CommonResponse
	AlibabaInteractMediaArtworkAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractMediaArtworkAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractMediaArtworkAPIResponseModel).Reset()
}

// AlibabaInteractMediaArtworkAPIResponseModel is 原图相关鉴权接口 成功返回结果
type AlibabaInteractMediaArtworkAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_media_artwork_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractMediaArtworkAPIResponseModel) Reset() {
	m.RequestId = ""
}

var poolAlibabaInteractMediaArtworkAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractMediaArtworkAPIResponse)
	},
}

// GetAlibabaInteractMediaArtworkAPIResponse 从 sync.Pool 获取 AlibabaInteractMediaArtworkAPIResponse
func GetAlibabaInteractMediaArtworkAPIResponse() *AlibabaInteractMediaArtworkAPIResponse {
	return poolAlibabaInteractMediaArtworkAPIResponse.Get().(*AlibabaInteractMediaArtworkAPIResponse)
}

// ReleaseAlibabaInteractMediaArtworkAPIResponse 将 AlibabaInteractMediaArtworkAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractMediaArtworkAPIResponse(v *AlibabaInteractMediaArtworkAPIResponse) {
	v.Reset()
	poolAlibabaInteractMediaArtworkAPIResponse.Put(v)
}
