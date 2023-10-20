package aedata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedata"
)

// Aliexpressaffiliateorderget AE流量订单详情获取API
// aliexpress.affiliate.order.get
//
// 联盟推广订单效果获取API
func Aliexpressaffiliateorderget(clt *core.SDKClient, req *aedata.AliexpressaffiliateordergetAPIRequest, session string) (*aedata.AliexpressaffiliateordergetAPIResponse, error) {
	var resp aedata.AliexpressaffiliateordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
