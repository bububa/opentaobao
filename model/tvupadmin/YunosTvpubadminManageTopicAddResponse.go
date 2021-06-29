package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增专题 API返回值 
yunos.tvpubadmin.manage.topic.add

专题新增
*/
type YunosTvpubadminManageTopicAddAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminManageTopicAddResponse
}

// 新增专题 成功返回结果
type YunosTvpubadminManageTopicAddResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`
}
