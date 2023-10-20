package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageTopicListAPIResponse 专题内容操作列表 API返回值
// yunos.tvpubadmin.manage.topic.list
//
// 获取外部可操作编辑的专题列表
type YunosTvpubadminManageTopicListAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminManageTopicListAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminManageTopicListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminManageTopicListAPIResponseModel).Reset()
}

// YunosTvpubadminManageTopicListAPIResponseModel is 专题内容操作列表 成功返回结果
type YunosTvpubadminManageTopicListAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminManageTopicListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminManageTopicListAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageTopicListAPIResponse)
	},
}

// GetYunosTvpubadminManageTopicListAPIResponse 从 sync.Pool 获取 YunosTvpubadminManageTopicListAPIResponse
func GetYunosTvpubadminManageTopicListAPIResponse() *YunosTvpubadminManageTopicListAPIResponse {
	return poolYunosTvpubadminManageTopicListAPIResponse.Get().(*YunosTvpubadminManageTopicListAPIResponse)
}

// ReleaseYunosTvpubadminManageTopicListAPIResponse 将 YunosTvpubadminManageTopicListAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminManageTopicListAPIResponse(v *YunosTvpubadminManageTopicListAPIResponse) {
	v.Reset()
	poolYunosTvpubadminManageTopicListAPIResponse.Put(v)
}
