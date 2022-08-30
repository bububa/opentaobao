package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsCommissionorderListbyindex 联盟订单分页查询
// aliexpress.ds.commissionorder.listbyindex
//
// 联盟订单分页查询
func AliexpressDsCommissionorderListbyindex(clt *core.SDKClient, req *aedropshiper.AliexpressDsCommissionorderListbyindexAPIRequest, session string) (*aedropshiper.AliexpressDsCommissionorderListbyindexAPIResponse, error) {
	var resp aedropshiper.AliexpressDsCommissionorderListbyindexAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
