package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdTargetTagEstimateUvAPIRequest
标签人群预估 API请求
alibaba.scbp.ad.target.tag.estimate.uv

标签人群预估 */
type AlibabaScbpAdTargetTagEstimateUvAPIRequest struct {
	model.Params
	// 请求参数
	_tagEstimateOperation *TagEstimateOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
