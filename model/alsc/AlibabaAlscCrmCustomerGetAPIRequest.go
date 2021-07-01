package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCustomerGetAPIRequest
查询顾客详情 API请求
alibaba.alsc.crm.customer.get

查询顾客详情 */
type AlibabaAlscCrmCustomerGetAPIRequest struct {
	model.Params
	// 顾客详情查询条件
	_paramCustomerIdQueryOpenReq *CustomerIdQueryOpenReq
}

// New
