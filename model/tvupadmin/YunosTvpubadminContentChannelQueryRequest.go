package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松影视频道查询 APIRequest
yunos.tvpubadmin.content.channel.query

迎客松影视频道查询
*/
type YunosTvpubadminContentChannelQueryRequest struct {
    model.Params

    // ChannelAuditQueryBO
    channelAuditQuery   string 

}

func NewYunosTvpubadminContentChannelQueryRequest() *YunosTvpubadminContentChannelQueryRequest{
    return &YunosTvpubadminContentChannelQueryRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentChannelQueryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.channel.query"
}

func (r YunosTvpubadminContentChannelQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentChannelQueryRequest) SetChannelAuditQuery(channelAuditQuery string) error {
    r.channelAuditQuery = channelAuditQuery
    r.Set("channel_audit_query", channelAuditQuery)
    return nil
}

func (r YunosTvpubadminContentChannelQueryRequest) GetChannelAuditQuery() string {
    return r.channelAuditQuery
}

