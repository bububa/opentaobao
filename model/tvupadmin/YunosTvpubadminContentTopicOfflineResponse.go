package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松专题下线 APIResponse
yunos.tvpubadmin.content.topic.offline

迎客松专题下线
*/
type YunosTvpubadminContentTopicOfflineAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentTopicOfflineResponse
}

type YunosTvpubadminContentTopicOfflineResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_topic_offline_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 下线操作结果
    
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`

    
}
