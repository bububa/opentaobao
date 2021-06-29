package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
专题内容操作列表 APIResponse
yunos.tvpubadmin.manage.topic.list

获取外部可操作编辑的专题列表
*/
type YunosTvpubadminManageTopicListAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminManageTopicListResponse
}

type YunosTvpubadminManageTopicListResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
