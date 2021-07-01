package tttm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunIndustryTttmOrderQueryAPIRequest
天天特卖数字工厂订单获取 API请求
aliyun.industry.tttm.order.query

获取阿里云数字工厂内天天特卖业务的订单 */
type AliyunIndustryTttmOrderQueryAPIRequest struct {
	model.Params
	// 订单号
	_orderId string
	// 外部采购单号
	_externalId string
}

// New
