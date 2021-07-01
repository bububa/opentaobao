package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentChannelOfflineAPIRequest
迎客松影视频道下线 API请求
yunos.tvpubadmin.content.channel.offline

迎客松影视频道下线 */
type YunosTvpubadminContentChannelOfflineAPIRequest struct {
	model.Params
	// id
	_id int64
}

// New
