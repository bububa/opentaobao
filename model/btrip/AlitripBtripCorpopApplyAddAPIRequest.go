package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCorpopApplyAddAPIRequest
【商旅】isv添加审批单 API请求
alitrip.btrip.corpop.apply.add

【商旅】isv添加审批单 */
type AlitripBtripCorpopApplyAddAPIRequest struct {
	model.Params
	// 请求参数
	_rq *OpenApiApplyRq
}

// New
