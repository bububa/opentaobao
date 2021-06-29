package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松影视频道查询 API请求
yunos.tvpubadmin.content.channel.query

迎客松影视频道查询
*/
type YunosTvpubadminContentChannelQueryRequest struct {
    model.Params
    // ChannelAuditQueryBO
    _channelAuditQuery   string
}

// 初始化YunosTvpubadminContentChannelQueryRequest对象
func NewYunosTvpubadminContentChannelQueryRequest() *YunosTvpubadminContentChannelQueryRequest{
    return &YunosTvpubadminContentChannelQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChannelQueryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.channel.query"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChannelQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelAuditQuery Setter
// ChannelAuditQueryBO
func (r *YunosTvpubadminContentChannelQueryRequest) SetChannelAuditQuery(_channelAuditQuery string) error {
    r._channelAuditQuery = _channelAuditQuery
    r.Set("channel_audit_query", _channelAuditQuery)
    return nil
}

// ChannelAuditQuery Getter
func (r YunosTvpubadminContentChannelQueryRequest) GetChannelAuditQuery() string {
    return r._channelAuditQuery
}
