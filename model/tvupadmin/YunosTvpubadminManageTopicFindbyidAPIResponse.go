package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageTopicFindbyidAPIResponse 根据id获取专题信息 API返回值
// yunos.tvpubadmin.manage.topic.findbyid
//
// 根据id获取专题信息
type YunosTvpubadminManageTopicFindbyidAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminManageTopicFindbyidAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminManageTopicFindbyidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminManageTopicFindbyidAPIResponseModel).Reset()
}

// YunosTvpubadminManageTopicFindbyidAPIResponseModel is 根据id获取专题信息 成功返回结果
type YunosTvpubadminManageTopicFindbyidAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_findbyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminManageTopicFindbyidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminManageTopicFindbyidAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageTopicFindbyidAPIResponse)
	},
}

// GetYunosTvpubadminManageTopicFindbyidAPIResponse 从 sync.Pool 获取 YunosTvpubadminManageTopicFindbyidAPIResponse
func GetYunosTvpubadminManageTopicFindbyidAPIResponse() *YunosTvpubadminManageTopicFindbyidAPIResponse {
	return poolYunosTvpubadminManageTopicFindbyidAPIResponse.Get().(*YunosTvpubadminManageTopicFindbyidAPIResponse)
}

// ReleaseYunosTvpubadminManageTopicFindbyidAPIResponse 将 YunosTvpubadminManageTopicFindbyidAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminManageTopicFindbyidAPIResponse(v *YunosTvpubadminManageTopicFindbyidAPIResponse) {
	v.Reset()
	poolYunosTvpubadminManageTopicFindbyidAPIResponse.Put(v)
}
