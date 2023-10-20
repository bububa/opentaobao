package taotv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvVideoPlaylistOttnavGetAPIResponse ott播单 API返回值
// taobao.taotv.video.playlist.ottnav.get
//
// 根据聚焦播单ID拿到下面播单视频，根据左侧播单ID列表批量拿到播单信息
type TaobaoTaotvVideoPlaylistOttnavGetAPIResponse struct {
	model.CommonResponse
	TaobaoTaotvVideoPlaylistOttnavGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTaotvVideoPlaylistOttnavGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTaotvVideoPlaylistOttnavGetAPIResponseModel).Reset()
}

// TaobaoTaotvVideoPlaylistOttnavGetAPIResponseModel is ott播单 成功返回结果
type TaobaoTaotvVideoPlaylistOttnavGetAPIResponseModel struct {
	XMLName xml.Name `xml:"taotv_video_playlist_ottnav_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoTaotvVideoPlaylistOttnavGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTaotvVideoPlaylistOttnavGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTaotvVideoPlaylistOttnavGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvVideoPlaylistOttnavGetAPIResponse)
	},
}

// GetTaobaoTaotvVideoPlaylistOttnavGetAPIResponse 从 sync.Pool 获取 TaobaoTaotvVideoPlaylistOttnavGetAPIResponse
func GetTaobaoTaotvVideoPlaylistOttnavGetAPIResponse() *TaobaoTaotvVideoPlaylistOttnavGetAPIResponse {
	return poolTaobaoTaotvVideoPlaylistOttnavGetAPIResponse.Get().(*TaobaoTaotvVideoPlaylistOttnavGetAPIResponse)
}

// ReleaseTaobaoTaotvVideoPlaylistOttnavGetAPIResponse 将 TaobaoTaotvVideoPlaylistOttnavGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTaotvVideoPlaylistOttnavGetAPIResponse(v *TaobaoTaotvVideoPlaylistOttnavGetAPIResponse) {
	v.Reset()
	poolTaobaoTaotvVideoPlaylistOttnavGetAPIResponse.Put(v)
}
