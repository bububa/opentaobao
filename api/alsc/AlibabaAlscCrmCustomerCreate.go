package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCustomerCreate 创建顾客
// alibaba.alsc.crm.customer.create
//
// 开放本地生活创建顾客功能
func AlibabaAlscCrmCustomerCreate(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCustomerCreateAPIRequest, resp *alsc.AlibabaAlscCrmCustomerCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
