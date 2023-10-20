package xiamicontent

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentSongsInfoQueryAPIResponse 搜索歌曲列表 API返回值
// xiami.content.songs.info.query
//
// 多维度查询歌曲列表
type XiamiContentSongsInfoQueryAPIResponse struct {
	model.CommonResponse
	XiamiContentSongsInfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *XiamiContentSongsInfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.XiamiContentSongsInfoQueryAPIResponseModel).Reset()
}

// XiamiContentSongsInfoQueryAPIResponseModel is 搜索歌曲列表 成功返回结果
type XiamiContentSongsInfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_songs_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的歌曲信息
	Songs *Page `json:"songs,omitempty" xml:"songs,omitempty"`
	// 系统自动生成
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *XiamiContentSongsInfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Songs = nil
	m.ResultCode = nil
}

var poolXiamiContentSongsInfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(XiamiContentSongsInfoQueryAPIResponse)
	},
}

// GetXiamiContentSongsInfoQueryAPIResponse 从 sync.Pool 获取 XiamiContentSongsInfoQueryAPIResponse
func GetXiamiContentSongsInfoQueryAPIResponse() *XiamiContentSongsInfoQueryAPIResponse {
	return poolXiamiContentSongsInfoQueryAPIResponse.Get().(*XiamiContentSongsInfoQueryAPIResponse)
}

// ReleaseXiamiContentSongsInfoQueryAPIResponse 将 XiamiContentSongsInfoQueryAPIResponse 保存到 sync.Pool
func ReleaseXiamiContentSongsInfoQueryAPIResponse(v *XiamiContentSongsInfoQueryAPIResponse) {
	v.Reset()
	poolXiamiContentSongsInfoQueryAPIResponse.Put(v)
}
