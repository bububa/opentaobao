package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCardBindcustomerAPIRequest
卡号绑定顾客 API请求
alibaba.alsc.crm.card.bindcustomer

为卡号绑定顾客 */
type AlibabaAlscCrmCardBindcustomerAPIRequest struct {
	model.Params
	// 请求参数
	_paramBindCustomerOpenReq *BindCustomerOpenReq
}

// New
