package aedata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAffiliateOrderListbyindexAPIRequest
AE联盟推广者订单查询接口-按游标索引查询 API请求
aliexpress.affiliate.order.listbyindex

AE联盟推广者订单按游标查询接口 */
type AliexpressAffiliateOrderListbyindexAPIRequest struct {
	model.Params
	// 开始时间
	_startTime string
	// 查询索引开始值：若不传，则只能查第一页
	_startQueryIndexId string
	// 结束时间
	_endTime string
	// 订单状态:Payment Completed,Buyer Confirmed Receipt
	_status string
	// 每页记录数
	_pageSize int64
	// 返回的字段信息
	_fields string
	// 安全签名
	_appSignature string
}

// New
