package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdGroupCountAdGroupAPIRequest
统计adgroup数量 API请求
alibaba.scbp.ad.group.count.ad.group

统计adgroup数量 */
type AlibabaScbpAdGroupCountAdGroupAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 查询条件
	_adGroupQuery *AdGroupQueryDto
	// 用户信息
	_topContext *TopContextDto
}

// New
