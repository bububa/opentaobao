package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCorpopApplyGetAPIRequest
【商旅】查询审批单 API请求
alitrip.btrip.corpop.apply.get

【商旅】查询审批单 */
type AlitripBtripCorpopApplyGetAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenIsvSearchRq
}

// New
