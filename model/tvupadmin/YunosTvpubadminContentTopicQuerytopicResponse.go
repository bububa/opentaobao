package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松专题查询 API返回值 
yunos.tvpubadmin.content.topic.querytopic

迎客松专题查询
*/
type YunosTvpubadminContentTopicQuerytopicAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentTopicQuerytopicResponse
}

// 迎客松专题查询 成功返回结果
type YunosTvpubadminContentTopicQuerytopicResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_topic_querytopic_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 专题列表
    Object   string `json:"object,omitempty" xml:"object,omitempty"`
}
