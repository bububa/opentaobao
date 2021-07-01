package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdCampaignDeleteAPIRequest
删除计划 API请求
alibaba.scbp.ad.campaign.delete

删除计划 */
type AlibabaScbpAdCampaignDeleteAPIRequest struct {
	model.Params
	// 操作对象
	_batchOperation *CampaignBatchOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
