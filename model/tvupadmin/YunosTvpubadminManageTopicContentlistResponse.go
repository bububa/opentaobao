package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查看专题内容列表 APIResponse
yunos.tvpubadmin.manage.topic.contentlist

查看专题内容列表
*/
type YunosTvpubadminManageTopicContentlistAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminManageTopicContentlistResponse
}

type YunosTvpubadminManageTopicContentlistResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_contentlist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
