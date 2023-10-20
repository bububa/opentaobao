package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentTopicQuerytopicAPIResponse 迎客松专题查询 API返回值
// yunos.tvpubadmin.content.topic.querytopic
//
// 迎客松专题查询
type YunosTvpubadminContentTopicQuerytopicAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentTopicQuerytopicAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentTopicQuerytopicAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentTopicQuerytopicAPIResponseModel).Reset()
}

// YunosTvpubadminContentTopicQuerytopicAPIResponseModel is 迎客松专题查询 成功返回结果
type YunosTvpubadminContentTopicQuerytopicAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_topic_querytopic_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 专题列表
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentTopicQuerytopicAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminContentTopicQuerytopicAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentTopicQuerytopicAPIResponse)
	},
}

// GetYunosTvpubadminContentTopicQuerytopicAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentTopicQuerytopicAPIResponse
func GetYunosTvpubadminContentTopicQuerytopicAPIResponse() *YunosTvpubadminContentTopicQuerytopicAPIResponse {
	return poolYunosTvpubadminContentTopicQuerytopicAPIResponse.Get().(*YunosTvpubadminContentTopicQuerytopicAPIResponse)
}

// ReleaseYunosTvpubadminContentTopicQuerytopicAPIResponse 将 YunosTvpubadminContentTopicQuerytopicAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentTopicQuerytopicAPIResponse(v *YunosTvpubadminContentTopicQuerytopicAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentTopicQuerytopicAPIResponse.Put(v)
}
