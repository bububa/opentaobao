package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageTopicAddAPIResponse 新增专题 API返回值
// yunos.tvpubadmin.manage.topic.add
//
// 专题新增
type YunosTvpubadminManageTopicAddAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminManageTopicAddAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminManageTopicAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminManageTopicAddAPIResponseModel).Reset()
}

// YunosTvpubadminManageTopicAddAPIResponseModel is 新增专题 成功返回结果
type YunosTvpubadminManageTopicAddAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *YunosTvpubadminManageTopicAddTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminManageTopicAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvpubadminManageTopicAddAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageTopicAddAPIResponse)
	},
}

// GetYunosTvpubadminManageTopicAddAPIResponse 从 sync.Pool 获取 YunosTvpubadminManageTopicAddAPIResponse
func GetYunosTvpubadminManageTopicAddAPIResponse() *YunosTvpubadminManageTopicAddAPIResponse {
	return poolYunosTvpubadminManageTopicAddAPIResponse.Get().(*YunosTvpubadminManageTopicAddAPIResponse)
}

// ReleaseYunosTvpubadminManageTopicAddAPIResponse 将 YunosTvpubadminManageTopicAddAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminManageTopicAddAPIResponse(v *YunosTvpubadminManageTopicAddAPIResponse) {
	v.Reset()
	poolYunosTvpubadminManageTopicAddAPIResponse.Put(v)
}
