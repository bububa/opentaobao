package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCharityUseractionSyncAPIRequest
用户公益行为同步 API请求
alibaba.charity.useraction.sync

外部公益活动，用户公益行为同步 */
type AlibabaCharityUseractionSyncAPIRequest struct {
	model.Params
	// 用户公益行为
	_channelUserActionDto *ChannelUserActionDto
}

// New
