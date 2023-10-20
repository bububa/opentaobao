package taotv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvVideoPlaylistAllAPIResponse 获取播单列表 API返回值
// taobao.taotv.video.playlist.all
//
// 根据牌照和视频源等获取播单列表
type TaobaoTaotvVideoPlaylistAllAPIResponse struct {
	model.CommonResponse
	TaobaoTaotvVideoPlaylistAllAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTaotvVideoPlaylistAllAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTaotvVideoPlaylistAllAPIResponseModel).Reset()
}

// TaobaoTaotvVideoPlaylistAllAPIResponseModel is 获取播单列表 成功返回结果
type TaobaoTaotvVideoPlaylistAllAPIResponseModel struct {
	XMLName xml.Name `xml:"taotv_video_playlist_all_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoTaotvVideoPlaylistAllResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTaotvVideoPlaylistAllAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTaotvVideoPlaylistAllAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvVideoPlaylistAllAPIResponse)
	},
}

// GetTaobaoTaotvVideoPlaylistAllAPIResponse 从 sync.Pool 获取 TaobaoTaotvVideoPlaylistAllAPIResponse
func GetTaobaoTaotvVideoPlaylistAllAPIResponse() *TaobaoTaotvVideoPlaylistAllAPIResponse {
	return poolTaobaoTaotvVideoPlaylistAllAPIResponse.Get().(*TaobaoTaotvVideoPlaylistAllAPIResponse)
}

// ReleaseTaobaoTaotvVideoPlaylistAllAPIResponse 将 TaobaoTaotvVideoPlaylistAllAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTaotvVideoPlaylistAllAPIResponse(v *TaobaoTaotvVideoPlaylistAllAPIResponse) {
	v.Reset()
	poolTaobaoTaotvVideoPlaylistAllAPIResponse.Put(v)
}
