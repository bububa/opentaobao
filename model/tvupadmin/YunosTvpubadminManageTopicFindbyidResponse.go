package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据id获取专题信息 APIResponse
yunos.tvpubadmin.manage.topic.findbyid

根据id获取专题信息
*/
type YunosTvpubadminManageTopicFindbyidAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminManageTopicFindbyidResponse
}

type YunosTvpubadminManageTopicFindbyidResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_findbyid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
