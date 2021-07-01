package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleTransferpayQueryAPIRequest
闲鱼转账结果查询 API请求
alibaba.idle.transferpay.query

商家业务 转账支付的结果查询 */
type AlibabaIdleTransferpayQueryAPIRequest struct {
	model.Params
	// 入参
	_param *PayQueryRequest
}

// New
