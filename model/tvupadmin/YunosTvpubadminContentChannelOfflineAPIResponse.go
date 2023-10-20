package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChannelOfflineAPIResponse 迎客松影视频道下线 API返回值
// yunos.tvpubadmin.content.channel.offline
//
// 迎客松影视频道下线
type YunosTvpubadminContentChannelOfflineAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentChannelOfflineAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentChannelOfflineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentChannelOfflineAPIResponseModel).Reset()
}

// YunosTvpubadminContentChannelOfflineAPIResponseModel is 迎客松影视频道下线 成功返回结果
type YunosTvpubadminContentChannelOfflineAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_channel_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 下线影视频道结果
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentChannelOfflineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = false
}

var poolYunosTvpubadminContentChannelOfflineAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentChannelOfflineAPIResponse)
	},
}

// GetYunosTvpubadminContentChannelOfflineAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentChannelOfflineAPIResponse
func GetYunosTvpubadminContentChannelOfflineAPIResponse() *YunosTvpubadminContentChannelOfflineAPIResponse {
	return poolYunosTvpubadminContentChannelOfflineAPIResponse.Get().(*YunosTvpubadminContentChannelOfflineAPIResponse)
}

// ReleaseYunosTvpubadminContentChannelOfflineAPIResponse 将 YunosTvpubadminContentChannelOfflineAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentChannelOfflineAPIResponse(v *YunosTvpubadminContentChannelOfflineAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentChannelOfflineAPIResponse.Put(v)
}
