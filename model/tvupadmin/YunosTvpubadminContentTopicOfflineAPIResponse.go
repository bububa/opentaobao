package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentTopicOfflineAPIResponse 迎客松专题下线 API返回值
// yunos.tvpubadmin.content.topic.offline
//
// 迎客松专题下线
type YunosTvpubadminContentTopicOfflineAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentTopicOfflineAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentTopicOfflineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentTopicOfflineAPIResponseModel).Reset()
}

// YunosTvpubadminContentTopicOfflineAPIResponseModel is 迎客松专题下线 成功返回结果
type YunosTvpubadminContentTopicOfflineAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_topic_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 下线操作结果
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentTopicOfflineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = false
}

var poolYunosTvpubadminContentTopicOfflineAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentTopicOfflineAPIResponse)
	},
}

// GetYunosTvpubadminContentTopicOfflineAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentTopicOfflineAPIResponse
func GetYunosTvpubadminContentTopicOfflineAPIResponse() *YunosTvpubadminContentTopicOfflineAPIResponse {
	return poolYunosTvpubadminContentTopicOfflineAPIResponse.Get().(*YunosTvpubadminContentTopicOfflineAPIResponse)
}

// ReleaseYunosTvpubadminContentTopicOfflineAPIResponse 将 YunosTvpubadminContentTopicOfflineAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentTopicOfflineAPIResponse(v *YunosTvpubadminContentTopicOfflineAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentTopicOfflineAPIResponse.Put(v)
}
