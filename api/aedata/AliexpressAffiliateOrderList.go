package aedata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedata"
)

// AliexpressAffiliateOrderList AE推广者订单批量获取接口
// aliexpress.affiliate.order.list
//
// AE联盟推广者订单分页查询接口
func AliexpressAffiliateOrderList(clt *core.SDKClient, req *aedata.AliexpressAffiliateOrderListAPIRequest, session string) (*aedata.AliexpressAffiliateOrderListAPIResponse, error) {
	var resp aedata.AliexpressAffiliateOrderListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
