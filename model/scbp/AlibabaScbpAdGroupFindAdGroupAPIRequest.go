package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdGroupFindAdGroupAPIRequest
查询推广组 API请求
alibaba.scbp.ad.group.find.ad.group

查询推广组 */
type AlibabaScbpAdGroupFindAdGroupAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 入参
	_adGroupQuery *AdGroupQueryDto
	// 用户信息
	_topContext *TopContextDto
}

// New
