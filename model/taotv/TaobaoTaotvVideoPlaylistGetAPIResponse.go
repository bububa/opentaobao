package taotv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvVideoPlaylistGetAPIResponse 根据频道ID获取频道下节目单以及当前播放 API返回值
// taobao.taotv.video.playlist.get
//
// 根据频道ID获取频道下节目单以及当前播放
type TaobaoTaotvVideoPlaylistGetAPIResponse struct {
	model.CommonResponse
	TaobaoTaotvVideoPlaylistGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTaotvVideoPlaylistGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTaotvVideoPlaylistGetAPIResponseModel).Reset()
}

// TaobaoTaotvVideoPlaylistGetAPIResponseModel is 根据频道ID获取频道下节目单以及当前播放 成功返回结果
type TaobaoTaotvVideoPlaylistGetAPIResponseModel struct {
	XMLName xml.Name `xml:"taotv_video_playlist_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoTaotvVideoPlaylistGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTaotvVideoPlaylistGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTaotvVideoPlaylistGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvVideoPlaylistGetAPIResponse)
	},
}

// GetTaobaoTaotvVideoPlaylistGetAPIResponse 从 sync.Pool 获取 TaobaoTaotvVideoPlaylistGetAPIResponse
func GetTaobaoTaotvVideoPlaylistGetAPIResponse() *TaobaoTaotvVideoPlaylistGetAPIResponse {
	return poolTaobaoTaotvVideoPlaylistGetAPIResponse.Get().(*TaobaoTaotvVideoPlaylistGetAPIResponse)
}

// ReleaseTaobaoTaotvVideoPlaylistGetAPIResponse 将 TaobaoTaotvVideoPlaylistGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTaotvVideoPlaylistGetAPIResponse(v *TaobaoTaotvVideoPlaylistGetAPIResponse) {
	v.Reset()
	poolTaobaoTaotvVideoPlaylistGetAPIResponse.Put(v)
}
