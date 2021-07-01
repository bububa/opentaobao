package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCustomerUpdateppwAPIRequest
修改支付密码 API请求
alibaba.alsc.crm.customer.updateppw

修改支付密码 */
type AlibabaAlscCrmCustomerUpdateppwAPIRequest struct {
	model.Params
	// 修改密码
	_updatePayPasswdReq *UpdatePayPasswdReq
}

// New
