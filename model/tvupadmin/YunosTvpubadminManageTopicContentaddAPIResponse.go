package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageTopicContentaddAPIResponse 专题新增内容 API返回值
// yunos.tvpubadmin.manage.topic.contentadd
//
// 专题新增内容
type YunosTvpubadminManageTopicContentaddAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminManageTopicContentaddAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminManageTopicContentaddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminManageTopicContentaddAPIResponseModel).Reset()
}

// YunosTvpubadminManageTopicContentaddAPIResponseModel is 专题新增内容 成功返回结果
type YunosTvpubadminManageTopicContentaddAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_contentadd_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *YunosTvpubadminManageTopicContentaddTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminManageTopicContentaddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvpubadminManageTopicContentaddAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageTopicContentaddAPIResponse)
	},
}

// GetYunosTvpubadminManageTopicContentaddAPIResponse 从 sync.Pool 获取 YunosTvpubadminManageTopicContentaddAPIResponse
func GetYunosTvpubadminManageTopicContentaddAPIResponse() *YunosTvpubadminManageTopicContentaddAPIResponse {
	return poolYunosTvpubadminManageTopicContentaddAPIResponse.Get().(*YunosTvpubadminManageTopicContentaddAPIResponse)
}

// ReleaseYunosTvpubadminManageTopicContentaddAPIResponse 将 YunosTvpubadminManageTopicContentaddAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminManageTopicContentaddAPIResponse(v *YunosTvpubadminManageTopicContentaddAPIResponse) {
	v.Reset()
	poolYunosTvpubadminManageTopicContentaddAPIResponse.Put(v)
}
