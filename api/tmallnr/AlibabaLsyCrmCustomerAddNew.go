package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmCustomerAddNew 导购域留资接口
// alibaba.lsy.crm.customer.add.new
//
// 导购域提供留资入口
func AlibabaLsyCrmCustomerAddNew(clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmCustomerAddNewAPIRequest, session string) (*tmallnr.AlibabaLsyCrmCustomerAddNewAPIResponse, error) {
	var resp tmallnr.AlibabaLsyCrmCustomerAddNewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
