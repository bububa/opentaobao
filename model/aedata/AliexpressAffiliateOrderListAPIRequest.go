package aedata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAffiliateOrderListAPIRequest
AE推广者订单批量获取接口 API请求
aliexpress.affiliate.order.list

AE联盟推广者订单分页查询接口 */
type AliexpressAffiliateOrderListAPIRequest struct {
	model.Params
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
	// 订单状态:Payment Completed,Buyer Confirmed Receipt
	_status string
	// 站点信息：global、ru_site、es_site、it_site
	_localeSite string
	// 页数
	_pageNo int64
	// 每页记录数
	_pageSize int64
	// 返回的字段信息
	_fields string
	// 安全签名
	_appSignature string
}

// New
