package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentChannelQueryAPIRequest
迎客松影视频道查询 API请求
yunos.tvpubadmin.content.channel.query

迎客松影视频道查询 */
type YunosTvpubadminContentChannelQueryAPIRequest struct {
	model.Params
	// ChannelAuditQueryBO
	_channelAuditQuery string
}

// New
