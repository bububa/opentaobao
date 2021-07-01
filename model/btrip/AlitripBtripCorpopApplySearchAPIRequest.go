package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCorpopApplySearchAPIRequest
【商旅】搜索审批单列表 API请求
alitrip.btrip.corpop.apply.search

【商旅】搜索审批单列表 */
type AlitripBtripCorpopApplySearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenIsvSearchRq
}

// New
