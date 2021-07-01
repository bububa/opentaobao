package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmOpenCustomerGetAPIRequest
查询会员资产 API请求
alibaba.alsc.crm.open.customer.get

查询会员身份，资产等 */
type AlibabaAlscCrmOpenCustomerGetAPIRequest struct {
	model.Params
	// 入参
	_paramCustomerGetOpenReq *CustomerGetOpenReq
}

// New
