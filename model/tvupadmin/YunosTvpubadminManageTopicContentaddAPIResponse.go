package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminManageTopicContentaddAPIResponse
专题新增内容 API返回值
yunos.tvpubadmin.manage.topic.contentadd

专题新增内容 */
type YunosTvpubadminManageTopicContentaddAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminManageTopicContentaddAPIResponseModel
}

// YunosTvpubadminManageTopicContentaddAPIResponseModel is 专题新增内容 成功返回结果
type YunosTvpubadminManageTopicContentaddAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_contentadd_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResult `json:"result,omitempty" xml:"result,omitempty"`
}
