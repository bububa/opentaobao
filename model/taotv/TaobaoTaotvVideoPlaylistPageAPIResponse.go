package taotv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvVideoPlaylistPageAPIResponse 分页获取所有播单 API返回值
// taobao.taotv.video.playlist.page
//
// 获取所有播单信息（分页）
type TaobaoTaotvVideoPlaylistPageAPIResponse struct {
	model.CommonResponse
	TaobaoTaotvVideoPlaylistPageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTaotvVideoPlaylistPageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTaotvVideoPlaylistPageAPIResponseModel).Reset()
}

// TaobaoTaotvVideoPlaylistPageAPIResponseModel is 分页获取所有播单 成功返回结果
type TaobaoTaotvVideoPlaylistPageAPIResponseModel struct {
	XMLName xml.Name `xml:"taotv_video_playlist_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoTaotvVideoPlaylistPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTaotvVideoPlaylistPageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTaotvVideoPlaylistPageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvVideoPlaylistPageAPIResponse)
	},
}

// GetTaobaoTaotvVideoPlaylistPageAPIResponse 从 sync.Pool 获取 TaobaoTaotvVideoPlaylistPageAPIResponse
func GetTaobaoTaotvVideoPlaylistPageAPIResponse() *TaobaoTaotvVideoPlaylistPageAPIResponse {
	return poolTaobaoTaotvVideoPlaylistPageAPIResponse.Get().(*TaobaoTaotvVideoPlaylistPageAPIResponse)
}

// ReleaseTaobaoTaotvVideoPlaylistPageAPIResponse 将 TaobaoTaotvVideoPlaylistPageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTaotvVideoPlaylistPageAPIResponse(v *TaobaoTaotvVideoPlaylistPageAPIResponse) {
	v.Reset()
	poolTaobaoTaotvVideoPlaylistPageAPIResponse.Put(v)
}
