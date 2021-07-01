package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsEfspUsercancelorderAPIRequest
用户取消订单 API请求
alibaba.alisports.efsp.usercancelorder

用户取消订单 */
type AlibabaAlisportsEfspUsercancelorderAPIRequest struct {
	model.Params
	// 订单编号
	_orderNo string
	// 用户支付宝ID
	_alipayId string
}

// New
