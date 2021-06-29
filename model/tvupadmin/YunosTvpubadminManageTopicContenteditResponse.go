package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
专题关联内容编辑 APIResponse
yunos.tvpubadmin.manage.topic.contentedit

编辑专题关联的内容
*/
type YunosTvpubadminManageTopicContenteditAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminManageTopicContenteditResponse
}

type YunosTvpubadminManageTopicContenteditResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_contentedit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 操作返回结果
    
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`

    
}
