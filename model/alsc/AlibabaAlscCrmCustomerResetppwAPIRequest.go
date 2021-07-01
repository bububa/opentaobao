package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCustomerResetppwAPIRequest
重置支付密码 API请求
alibaba.alsc.crm.customer.resetppw

重置支付密码 */
type AlibabaAlscCrmCustomerResetppwAPIRequest struct {
	model.Params
	// 系统自动生成
	_resetPayPwdRequest *ResetPayPasswdOpenReq
}

// New
