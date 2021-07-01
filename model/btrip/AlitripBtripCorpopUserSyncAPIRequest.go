package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCorpopUserSyncAPIRequest
外部人员同步 API请求
alitrip.btrip.corpop.user.sync

同步外部平台用户信息至商旅内部 */
type AlitripBtripCorpopUserSyncAPIRequest struct {
	model.Params
	// 人员同步请求
	_rq *BtripUserSyncRq
}

// New
