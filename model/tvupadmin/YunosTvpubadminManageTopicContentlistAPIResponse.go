package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageTopicContentlistAPIResponse 查看专题内容列表 API返回值
// yunos.tvpubadmin.manage.topic.contentlist
//
// 查看专题内容列表
type YunosTvpubadminManageTopicContentlistAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminManageTopicContentlistAPIResponseModel
}

// YunosTvpubadminManageTopicContentlistAPIResponseModel is 查看专题内容列表 成功返回结果
type YunosTvpubadminManageTopicContentlistAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_contentlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}
