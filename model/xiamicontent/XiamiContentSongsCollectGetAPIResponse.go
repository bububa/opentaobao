package xiamicontent

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentSongsCollectGetAPIResponse 获取歌单详情接口 API返回值
// xiami.content.songs.collect.get
//
// 根据歌单id，获取歌单详情
type XiamiContentSongsCollectGetAPIResponse struct {
	model.CommonResponse
	XiamiContentSongsCollectGetAPIResponseModel
}

// Reset 清空结构体
func (m *XiamiContentSongsCollectGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.XiamiContentSongsCollectGetAPIResponseModel).Reset()
}

// XiamiContentSongsCollectGetAPIResponseModel is 获取歌单详情接口 成功返回结果
type XiamiContentSongsCollectGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_songs_collect_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的歌曲信息
	Songs *Page `json:"songs,omitempty" xml:"songs,omitempty"`
	// 系统自动生成
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 歌单详情
	Collect *CollectDto `json:"collect,omitempty" xml:"collect,omitempty"`
}

// Reset 清空结构体
func (m *XiamiContentSongsCollectGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Songs = nil
	m.ResultCode = nil
	m.Collect = nil
}

var poolXiamiContentSongsCollectGetAPIResponse = sync.Pool{
	New: func() any {
		return new(XiamiContentSongsCollectGetAPIResponse)
	},
}

// GetXiamiContentSongsCollectGetAPIResponse 从 sync.Pool 获取 XiamiContentSongsCollectGetAPIResponse
func GetXiamiContentSongsCollectGetAPIResponse() *XiamiContentSongsCollectGetAPIResponse {
	return poolXiamiContentSongsCollectGetAPIResponse.Get().(*XiamiContentSongsCollectGetAPIResponse)
}

// ReleaseXiamiContentSongsCollectGetAPIResponse 将 XiamiContentSongsCollectGetAPIResponse 保存到 sync.Pool
func ReleaseXiamiContentSongsCollectGetAPIResponse(v *XiamiContentSongsCollectGetAPIResponse) {
	v.Reset()
	poolXiamiContentSongsCollectGetAPIResponse.Put(v)
}
