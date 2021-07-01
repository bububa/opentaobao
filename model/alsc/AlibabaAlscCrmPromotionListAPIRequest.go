package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmPromotionListAPIRequest
获取促销规则列表 API请求
alibaba.alsc.crm.promotion.list

获取品牌的促销规则列表 */
type AlibabaAlscCrmPromotionListAPIRequest struct {
	model.Params
	// 获取促销规则请求参数
	_promotionFacadeOpenReq *PromotionFacadeOpenReq
}

// New
