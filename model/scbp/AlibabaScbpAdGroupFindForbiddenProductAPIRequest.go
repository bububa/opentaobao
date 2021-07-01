package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdGroupFindForbiddenProductAPIRequest
查询屏蔽品 API请求
alibaba.scbp.ad.group.find.forbidden.product

查询屏蔽品 */
type AlibabaScbpAdGroupFindForbiddenProductAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 用户信息
	_topContext *TopContextDto
}

// New
