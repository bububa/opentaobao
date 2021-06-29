package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松专题查询 APIResponse
yunos.tvpubadmin.content.topic.querytopic

迎客松专题查询
*/
type YunosTvpubadminContentTopicQuerytopicAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentTopicQuerytopicResponse
}

type YunosTvpubadminContentTopicQuerytopicResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_topic_querytopic_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 专题列表
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
