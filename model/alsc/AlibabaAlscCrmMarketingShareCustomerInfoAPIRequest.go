package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest
查询分享营销客户领券信息 API请求
alibaba.alsc.crm.marketing.share.customer.info

查询分享营销活动的客户领券信息 */
type AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest struct {
	model.Params
	// 活动id
	_activityId string
	// 会员id
	_customerId string
	// 品牌id(brandId和outerBrandId必传其一)
	_brandId string
	// 操作人
	_operatorId string
	// 操作人姓名
	_operatorName string
	// 外部品牌id
	_outBrandId string
	// 外部门店id
	_outShopId string
	// 请求幂等id
	_requestId string
	// 门店id
	_shopId string
}

// New
