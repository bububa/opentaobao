package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentchannelqueryAPIRequest 迎客松影视频道查询 API请求
// yunos.tvpubadmin.content.channel.query
//
// 迎客松影视频道查询
type YunostvpubadmincontentchannelqueryAPIRequest struct {
	model.Params
	// ChannelAuditQueryBO
	_channelAuditQuery string
}

// NewYunostvpubadmincontentchannelqueryRequest 初始化YunostvpubadmincontentchannelqueryAPIRequest对象
func NewYunostvpubadmincontentchannelqueryRequest() *YunostvpubadmincontentchannelqueryAPIRequest {
	return &YunostvpubadmincontentchannelqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentchannelqueryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.channel.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentchannelqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentchannelqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelAuditQuery is ChannelAuditQuery Setter
// ChannelAuditQueryBO
func (r *YunostvpubadmincontentchannelqueryAPIRequest) SetChannelAuditQuery(_channelAuditQuery string) error {
	r._channelAuditQuery = _channelAuditQuery
	r.Set("channel_audit_query", _channelAuditQuery)
	return nil
}

// GetChannelAuditQuery ChannelAuditQuery Getter
func (r YunostvpubadmincontentchannelqueryAPIRequest) GetChannelAuditQuery() string {
	return r._channelAuditQuery
}
