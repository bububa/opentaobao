package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdCampaignEffectAPIRequest
定向推广-获取计划维度推广效果 API请求
alibaba.scbp.target.ad.campaign.effect

定向推广-获取计划维度推广效果 */
type AlibabaScbpTargetAdCampaignEffectAPIRequest struct {
	model.Params
	// 统计区间 只能为1 7 30
	_interval int64
	// 结束时间 当inteval=7或30的时候 不需要填写，当inteval=1时需要填写（开始结束时间区间不允许大于180天）
	_endDate string
	// 开始时间 当inteval=7或30的时候 不需要填写，当inteval=1时需要填写（开始结束时间区间不允许大于180天）
	_beginDate string
	// 当填写时，展示指定id的数据，不填写，则展示全部计划总数据
	_campaignId int64
}

// New
