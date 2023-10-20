package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsCommissionorderListbyindex 联盟订单分页查询
// aliexpress.ds.commissionorder.listbyindex
//
// 联盟订单分页查询
func AliexpressDsCommissionorderListbyindex(clt *core.SDKClient, req *aedropshiper.AliexpressDsCommissionorderListbyindexAPIRequest, resp *aedropshiper.AliexpressDsCommissionorderListbyindexAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
