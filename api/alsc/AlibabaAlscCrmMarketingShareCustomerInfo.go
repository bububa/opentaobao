package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmmarketingsharecustomerinfo 查询分享营销客户领券信息
// alibaba.alsc.crm.marketing.share.customer.info
//
// 查询分享营销活动的客户领券信息
func Alibabaalsccrmmarketingsharecustomerinfo(clt *core.SDKClient, req *alsc.AlibabaalsccrmmarketingsharecustomerinfoAPIRequest, session string) (*alsc.AlibabaalsccrmmarketingsharecustomerinfoAPIResponse, error) {
	var resp alsc.AlibabaalsccrmmarketingsharecustomerinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
