package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCorpopDepartSyncAPIRequest
外部部门同步 API请求
alitrip.btrip.corpop.depart.sync

同步外部平台部门信息至商旅内部 */
type AlitripBtripCorpopDepartSyncAPIRequest struct {
	model.Params
	// 同步部门请求
	_rq *BtripDepartSyncRq
}

// New
