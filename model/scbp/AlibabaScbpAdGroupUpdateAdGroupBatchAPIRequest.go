package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest
修改推广单元 API请求
alibaba.scbp.ad.group.update.ad.group.batch

修改推广单元 */
type AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 入参
	_adGroupBatchOperation *AdGroupBatchOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
