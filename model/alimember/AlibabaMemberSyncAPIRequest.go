package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMemberSyncAPIRequest
会员信息同步 API请求
alibaba.member.sync

会员信息同步 */
type AlibabaMemberSyncAPIRequest struct {
	model.Params
	// 会员同步信息
	_syncMember *SyncMemberDto
}

// New
