package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChannelQueryAPIRequest 迎客松影视频道查询 API请求
// yunos.tvpubadmin.content.channel.query
//
// 迎客松影视频道查询
type YunosTvpubadminContentChannelQueryAPIRequest struct {
	model.Params
	// ChannelAuditQueryBO
	_channelAuditQuery string
}

// NewYunosTvpubadminContentChannelQueryRequest 初始化YunosTvpubadminContentChannelQueryAPIRequest对象
func NewYunosTvpubadminContentChannelQueryRequest() *YunosTvpubadminContentChannelQueryAPIRequest {
	return &YunosTvpubadminContentChannelQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentChannelQueryAPIRequest) Reset() {
	r._channelAuditQuery = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChannelQueryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.channel.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChannelQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentChannelQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelAuditQuery is ChannelAuditQuery Setter
// ChannelAuditQueryBO
func (r *YunosTvpubadminContentChannelQueryAPIRequest) SetChannelAuditQuery(_channelAuditQuery string) error {
	r._channelAuditQuery = _channelAuditQuery
	r.Set("channel_audit_query", _channelAuditQuery)
	return nil
}

// GetChannelAuditQuery ChannelAuditQuery Getter
func (r YunosTvpubadminContentChannelQueryAPIRequest) GetChannelAuditQuery() string {
	return r._channelAuditQuery
}

var poolYunosTvpubadminContentChannelQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentChannelQueryRequest()
	},
}

// GetYunosTvpubadminContentChannelQueryRequest 从 sync.Pool 获取 YunosTvpubadminContentChannelQueryAPIRequest
func GetYunosTvpubadminContentChannelQueryAPIRequest() *YunosTvpubadminContentChannelQueryAPIRequest {
	return poolYunosTvpubadminContentChannelQueryAPIRequest.Get().(*YunosTvpubadminContentChannelQueryAPIRequest)
}

// ReleaseYunosTvpubadminContentChannelQueryAPIRequest 将 YunosTvpubadminContentChannelQueryAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentChannelQueryAPIRequest(v *YunosTvpubadminContentChannelQueryAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentChannelQueryAPIRequest.Put(v)
}
