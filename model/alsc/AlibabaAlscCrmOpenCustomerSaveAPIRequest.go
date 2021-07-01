package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmOpenCustomerSaveAPIRequest
保存和更新顾客 API请求
alibaba.alsc.crm.open.customer.save

用来保存顾客，如果已经存在的话，则更新顾客 */
type AlibabaAlscCrmOpenCustomerSaveAPIRequest struct {
	model.Params
	// 入参
	_paramCustomerSaveOpenReq *CustomerSaveOpenReq
}

// New
