package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCustomerCheckppwAPIRequest
校验支付密码 API请求
alibaba.alsc.crm.customer.checkppw

校验支付密码 */
type AlibabaAlscCrmCustomerCheckppwAPIRequest struct {
	model.Params
	// 请求参数
	_checkRequest *CheckPayPasswdReq
}

// New
