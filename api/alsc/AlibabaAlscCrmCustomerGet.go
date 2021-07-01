package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

/* AlibabaAlscCrmCustomerGet
查询顾客详情
alibaba.alsc.crm.customer.get

查询顾客详情 */
func AlibabaAlscCrmCustomerGet(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCustomerGetAPIRequest, session string) (*alsc.AlibabaAlscCrmCustomerGetAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmCustomerGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
