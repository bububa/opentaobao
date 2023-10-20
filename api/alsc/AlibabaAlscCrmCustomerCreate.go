package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCustomerCreate 创建顾客
// alibaba.alsc.crm.customer.create
//
// 开放本地生活创建顾客功能
func AlibabaAlscCrmCustomerCreate(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCustomerCreateAPIRequest, session string) (*alsc.AlibabaAlscCrmCustomerCreateAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmCustomerCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
