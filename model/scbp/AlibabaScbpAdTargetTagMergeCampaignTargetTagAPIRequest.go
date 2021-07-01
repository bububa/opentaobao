package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest
标签增删改 API请求
alibaba.scbp.ad.target.tag.merge.campaign.target.tag

标签增删改 */
type AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 标签数据，json格式。 最外层key：人群标签crowd/地域标签region、priceMode，第二层key: 增add、删del、改mod，第三层key：optionValue、bidRate、tagId  eg: 删除：{"crowd":{"del":[{"tagId":3595769030}]}}   修改：{"crowd":{"mod":[{"optionValue":"high_potential_order_user","bidRate":"151"}]}} 增加：{"crowd":{"add":[{"optionValue":"user_area_CA","bidRate":"133"}]}}
	_data string
	// 用户信息
	_topContext *TopContextDto
}

// New
