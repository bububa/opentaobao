package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsEfspUserplaceorderAPIRequest
用户完成支付同步订单 API请求
alibaba.alisports.efsp.userplaceorder

用户完成支付同步订单 */
type AlibabaAlisportsEfspUserplaceorderAPIRequest struct {
	model.Params
	// 青橙订单的json
	_orderJson string
}

// New
