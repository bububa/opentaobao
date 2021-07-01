package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdGroupCreateForbiddenProductAPIRequest
创建屏蔽品 API请求
alibaba.scbp.ad.group.create.forbidden.product

创建屏蔽品 */
type AlibabaScbpAdGroupCreateForbiddenProductAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 查询条件
	_forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
