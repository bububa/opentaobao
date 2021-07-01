package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCustomerCreateAPIRequest
创建顾客 API请求
alibaba.alsc.crm.customer.create

开放本地生活创建顾客功能 */
type AlibabaAlscCrmCustomerCreateAPIRequest struct {
	model.Params
	// 创建顾客参数
	_paramCustomerCreateOpenReq *CustomerCreateOpenReq
}

// New
