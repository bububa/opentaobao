package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest
创建推广单元 API请求
alibaba.scbp.ad.group.create.ad.group.batch

创建推广单元 */
type AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 入参
	_adGroupBatchOperation *AdGroupBatchOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
