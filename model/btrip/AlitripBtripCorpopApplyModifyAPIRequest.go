package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCorpopApplyModifyAPIRequest
【商旅】修改出差审批单（行程） API请求
alitrip.btrip.corpop.apply.modify

【商旅】修改出差审批单（行程） */
type AlitripBtripCorpopApplyModifyAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiApplyRq
}

// New
