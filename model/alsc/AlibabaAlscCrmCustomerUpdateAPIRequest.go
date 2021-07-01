package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCustomerUpdateAPIRequest
更新顾客信息 API请求
alibaba.alsc.crm.customer.update

更新顾客信息 */
type AlibabaAlscCrmCustomerUpdateAPIRequest struct {
	model.Params
	// 修改顾客参数
	_paramCustomerUpdateOpenReq *CustomerUpdateOpenReq
}

// New
