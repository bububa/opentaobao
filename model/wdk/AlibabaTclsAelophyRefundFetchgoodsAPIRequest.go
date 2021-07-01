package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyRefundFetchgoodsAPIRequest
saas 售后逆向 商户发起逆向取货 API请求
alibaba.tcls.aelophy.refund.fetchgoods

saas 售后逆向 商户发起逆向取货 */
type AlibabaTclsAelophyRefundFetchgoodsAPIRequest struct {
	model.Params
	// 经营店ID
	_storeId string
	// 外部订单ID
	_outOrderId string
	// 退款单ID
	_refundId string
	// 取货开始时间
	_fetchStartTime string
	// 取货结束时间
	_fetchEndTime string
	// 备注
	_remark string
	// 外部子订单列表
	_subRefundList []Subrefundlist
}

// New
