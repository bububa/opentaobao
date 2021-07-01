package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest
删除推广单元 API请求
alibaba.scbp.ad.group.delete.ad.group.batch

删除推广单元 */
type AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 请求参数
	_adGroupBatchOperation *AdGroupBatchOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
