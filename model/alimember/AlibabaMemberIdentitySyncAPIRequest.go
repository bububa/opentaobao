package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMemberIdentitySyncAPIRequest
会员身份信息同步 API请求
alibaba.member.identity.sync

会员身份信息同步 */
type AlibabaMemberIdentitySyncAPIRequest struct {
	model.Params
	// 会员身份同步信息
	_syncDto *SyncMemberIdentityDto
}

// New
