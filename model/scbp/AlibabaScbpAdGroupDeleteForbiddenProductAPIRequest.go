package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest
删除屏蔽品 API请求
alibaba.scbp.ad.group.delete.forbidden.product

删除屏蔽品 */
type AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 请求参数
	_forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
