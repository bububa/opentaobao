package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmCustomerAdd 私域导购添加活动留资入口
// alibaba.lsy.crm.customer.add
//
// 私域导购添加活动留资入口
func AlibabaLsyCrmCustomerAdd(clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmCustomerAddAPIRequest, resp *tmallnr.AlibabaLsyCrmCustomerAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
