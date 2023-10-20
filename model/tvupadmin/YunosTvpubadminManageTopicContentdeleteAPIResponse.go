package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageTopicContentdeleteAPIResponse 删除专题下内容 API返回值
// yunos.tvpubadmin.manage.topic.contentdelete
//
// 删除专题下内容信息
type YunosTvpubadminManageTopicContentdeleteAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminManageTopicContentdeleteAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminManageTopicContentdeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminManageTopicContentdeleteAPIResponseModel).Reset()
}

// YunosTvpubadminManageTopicContentdeleteAPIResponseModel is 删除专题下内容 成功返回结果
type YunosTvpubadminManageTopicContentdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_contentdelete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *YunosTvpubadminManageTopicContentdeleteTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminManageTopicContentdeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvpubadminManageTopicContentdeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageTopicContentdeleteAPIResponse)
	},
}

// GetYunosTvpubadminManageTopicContentdeleteAPIResponse 从 sync.Pool 获取 YunosTvpubadminManageTopicContentdeleteAPIResponse
func GetYunosTvpubadminManageTopicContentdeleteAPIResponse() *YunosTvpubadminManageTopicContentdeleteAPIResponse {
	return poolYunosTvpubadminManageTopicContentdeleteAPIResponse.Get().(*YunosTvpubadminManageTopicContentdeleteAPIResponse)
}

// ReleaseYunosTvpubadminManageTopicContentdeleteAPIResponse 将 YunosTvpubadminManageTopicContentdeleteAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminManageTopicContentdeleteAPIResponse(v *YunosTvpubadminManageTopicContentdeleteAPIResponse) {
	v.Reset()
	poolYunosTvpubadminManageTopicContentdeleteAPIResponse.Put(v)
}
