package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChannelQueryAPIResponse 迎客松影视频道查询 API返回值
// yunos.tvpubadmin.content.channel.query
//
// 迎客松影视频道查询
type YunosTvpubadminContentChannelQueryAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentChannelQueryAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentChannelQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentChannelQueryAPIResponseModel).Reset()
}

// YunosTvpubadminContentChannelQueryAPIResponseModel is 迎客松影视频道查询 成功返回结果
type YunosTvpubadminContentChannelQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_channel_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 影视频道列表
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentChannelQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminContentChannelQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentChannelQueryAPIResponse)
	},
}

// GetYunosTvpubadminContentChannelQueryAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentChannelQueryAPIResponse
func GetYunosTvpubadminContentChannelQueryAPIResponse() *YunosTvpubadminContentChannelQueryAPIResponse {
	return poolYunosTvpubadminContentChannelQueryAPIResponse.Get().(*YunosTvpubadminContentChannelQueryAPIResponse)
}

// ReleaseYunosTvpubadminContentChannelQueryAPIResponse 将 YunosTvpubadminContentChannelQueryAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentChannelQueryAPIResponse(v *YunosTvpubadminContentChannelQueryAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentChannelQueryAPIResponse.Put(v)
}
