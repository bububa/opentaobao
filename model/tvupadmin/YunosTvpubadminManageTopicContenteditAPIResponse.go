package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageTopicContenteditAPIResponse 专题关联内容编辑 API返回值
// yunos.tvpubadmin.manage.topic.contentedit
//
// 编辑专题关联的内容
type YunosTvpubadminManageTopicContenteditAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminManageTopicContenteditAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminManageTopicContenteditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminManageTopicContenteditAPIResponseModel).Reset()
}

// YunosTvpubadminManageTopicContenteditAPIResponseModel is 专题关联内容编辑 成功返回结果
type YunosTvpubadminManageTopicContenteditAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_contentedit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作返回结果
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminManageTopicContenteditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = false
}

var poolYunosTvpubadminManageTopicContenteditAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageTopicContenteditAPIResponse)
	},
}

// GetYunosTvpubadminManageTopicContenteditAPIResponse 从 sync.Pool 获取 YunosTvpubadminManageTopicContenteditAPIResponse
func GetYunosTvpubadminManageTopicContenteditAPIResponse() *YunosTvpubadminManageTopicContenteditAPIResponse {
	return poolYunosTvpubadminManageTopicContenteditAPIResponse.Get().(*YunosTvpubadminManageTopicContenteditAPIResponse)
}

// ReleaseYunosTvpubadminManageTopicContenteditAPIResponse 将 YunosTvpubadminManageTopicContenteditAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminManageTopicContenteditAPIResponse(v *YunosTvpubadminManageTopicContenteditAPIResponse) {
	v.Reset()
	poolYunosTvpubadminManageTopicContenteditAPIResponse.Put(v)
}
