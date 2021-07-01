package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest
查询标签数据 API请求
alibaba.scbp.ad.target.tag.find.campaign.target.tag

查询标签数据 */
type AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 请求参数
	_targetTagOperation *TargetTagOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// New
