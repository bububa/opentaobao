package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMusicSearchAPIResponse 对外音乐搜索服务 API返回值
// taobao.ailab.aicloud.top.music.search
//
// 供厂商获取音乐列表
type TaobaoAilabAicloudTopMusicSearchAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopMusicSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMusicSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopMusicSearchAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopMusicSearchAPIResponseModel is 对外音乐搜索服务 成功返回结果
type TaobaoAilabAicloudTopMusicSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_music_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMusicSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopMusicSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMusicSearchAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopMusicSearchAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopMusicSearchAPIResponse
func GetTaobaoAilabAicloudTopMusicSearchAPIResponse() *TaobaoAilabAicloudTopMusicSearchAPIResponse {
	return poolTaobaoAilabAicloudTopMusicSearchAPIResponse.Get().(*TaobaoAilabAicloudTopMusicSearchAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopMusicSearchAPIResponse 将 TaobaoAilabAicloudTopMusicSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopMusicSearchAPIResponse(v *TaobaoAilabAicloudTopMusicSearchAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopMusicSearchAPIResponse.Put(v)
}
