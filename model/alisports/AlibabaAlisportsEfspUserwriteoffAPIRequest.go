package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsEfspUserwriteoffAPIRequest
用户核销 API请求
alibaba.alisports.efsp.userwriteoff

用户核销 */
type AlibabaAlisportsEfspUserwriteoffAPIRequest struct {
	model.Params
	// 订单编号
	_orderNo string
	// 订单金额
	_sumAmount int64
	// 健身房Id
	_gymId string
	// 用户支付宝ID
	_alipayId string
	// 补助金额
	_subsidyAmount int64
}

// New
