package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieTracksSearchAPIResponse 查询音频 API返回值
// alibaba.ailabs.aligenie.tracks.search
//
// 搜索类目下的音频信息
type AlibabaAilabsAligenieTracksSearchAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieTracksSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieTracksSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsAligenieTracksSearchAPIResponseModel).Reset()
}

// AlibabaAilabsAligenieTracksSearchAPIResponseModel is 查询音频 成功返回结果
type AlibabaAilabsAligenieTracksSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_tracks_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieTracksSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsAligenieTracksSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsAligenieTracksSearchAPIResponse)
	},
}

// GetAlibabaAilabsAligenieTracksSearchAPIResponse 从 sync.Pool 获取 AlibabaAilabsAligenieTracksSearchAPIResponse
func GetAlibabaAilabsAligenieTracksSearchAPIResponse() *AlibabaAilabsAligenieTracksSearchAPIResponse {
	return poolAlibabaAilabsAligenieTracksSearchAPIResponse.Get().(*AlibabaAilabsAligenieTracksSearchAPIResponse)
}

// ReleaseAlibabaAilabsAligenieTracksSearchAPIResponse 将 AlibabaAilabsAligenieTracksSearchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsAligenieTracksSearchAPIResponse(v *AlibabaAilabsAligenieTracksSearchAPIResponse) {
	v.Reset()
	poolAlibabaAilabsAligenieTracksSearchAPIResponse.Put(v)
}
