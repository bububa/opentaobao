package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripXhotelChannelOrderMembertypeSyncAPIRequest
酒店分销渠道会员类型同步 API请求
alitrip.xhotel.channel.order.membertype.sync

酒店分销渠道会员类型同步 */
type AlitripXhotelChannelOrderMembertypeSyncAPIRequest struct {
	model.Params
	// 入参
	_channelSyncOrderMemberType *ChannelSyncOrderMemberType
}

// New
