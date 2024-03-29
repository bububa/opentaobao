package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageTopicEditAPIResponse 编辑专题 API返回值
// yunos.tvpubadmin.manage.topic.edit
//
// 编辑专题
type YunosTvpubadminManageTopicEditAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminManageTopicEditAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminManageTopicEditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminManageTopicEditAPIResponseModel).Reset()
}

// YunosTvpubadminManageTopicEditAPIResponseModel is 编辑专题 成功返回结果
type YunosTvpubadminManageTopicEditAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *YunosTvpubadminManageTopicEditTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminManageTopicEditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvpubadminManageTopicEditAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageTopicEditAPIResponse)
	},
}

// GetYunosTvpubadminManageTopicEditAPIResponse 从 sync.Pool 获取 YunosTvpubadminManageTopicEditAPIResponse
func GetYunosTvpubadminManageTopicEditAPIResponse() *YunosTvpubadminManageTopicEditAPIResponse {
	return poolYunosTvpubadminManageTopicEditAPIResponse.Get().(*YunosTvpubadminManageTopicEditAPIResponse)
}

// ReleaseYunosTvpubadminManageTopicEditAPIResponse 将 YunosTvpubadminManageTopicEditAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminManageTopicEditAPIResponse(v *YunosTvpubadminManageTopicEditAPIResponse) {
	v.Reset()
	poolYunosTvpubadminManageTopicEditAPIResponse.Put(v)
}
