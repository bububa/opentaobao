package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkOrderRefundGetAPIRequest
五道口订单退款按ID查询 API请求
alibaba.wdk.order.refund.get

按照退款ID或者五道口中台订单ID查询退款信息详情 */
type AlibabaWdkOrderRefundGetAPIRequest struct {
	model.Params
	// 五道口订单列表（子订单）
	_bizOrderIds []int64
	// 退款订单列表
	_refundIds []int64
	// 渠道来源 3：饿了么 4：盒马
	_orderFrom int64
	// 渠道店id
	_shopId string
	// 经营店id
	_storeId string
}

// New
