package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmCustomerAddNew 导购域留资接口
// alibaba.lsy.crm.customer.add.new
//
// 导购域提供留资入口
func AlibabaLsyCrmCustomerAddNew(ctx context.Context, clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmCustomerAddNewAPIRequest, resp *tmallnr.AlibabaLsyCrmCustomerAddNewAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
