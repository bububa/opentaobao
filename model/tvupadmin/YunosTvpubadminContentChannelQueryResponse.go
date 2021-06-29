package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松影视频道查询 APIResponse
yunos.tvpubadmin.content.channel.query

迎客松影视频道查询
*/
type YunosTvpubadminContentChannelQueryAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentChannelQueryResponse
}

type YunosTvpubadminContentChannelQueryResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_channel_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 影视频道列表
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
