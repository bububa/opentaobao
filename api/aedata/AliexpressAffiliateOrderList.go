package aedata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedata"
)

// Aliexpressaffiliateorderlist AE推广者订单批量获取接口
// aliexpress.affiliate.order.list
//
// AE联盟推广者订单分页查询接口
func Aliexpressaffiliateorderlist(clt *core.SDKClient, req *aedata.AliexpressaffiliateorderlistAPIRequest, session string) (*aedata.AliexpressaffiliateorderlistAPIResponse, error) {
	var resp aedata.AliexpressaffiliateorderlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
