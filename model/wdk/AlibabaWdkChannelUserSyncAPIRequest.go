package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkChannelUserSyncAPIRequest
会员同步 API请求
alibaba.wdk.channel.user.sync

会员同步 */
type AlibabaWdkChannelUserSyncAPIRequest struct {
	model.Params
	// 会员信息
	_userSyncInfo *UserSyncInfo
}

// New
