package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// Aliexpressdscommissionorderlistbyindex 联盟订单分页查询
// aliexpress.ds.commissionorder.listbyindex
//
// 联盟订单分页查询
func Aliexpressdscommissionorderlistbyindex(clt *core.SDKClient, req *aedropshiper.AliexpressdscommissionorderlistbyindexAPIRequest, session string) (*aedropshiper.AliexpressdscommissionorderlistbyindexAPIResponse, error) {
	var resp aedropshiper.AliexpressdscommissionorderlistbyindexAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
