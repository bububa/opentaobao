package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松影视频道下线 APIResponse
yunos.tvpubadmin.content.channel.offline

迎客松影视频道下线
*/
type YunosTvpubadminContentChannelOfflineAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentChannelOfflineResponse
}

type YunosTvpubadminContentChannelOfflineResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_channel_offline_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 下线影视频道结果
    
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`

    
}
