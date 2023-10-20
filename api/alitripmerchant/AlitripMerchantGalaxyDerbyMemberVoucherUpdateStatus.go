package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervoucherupdatestatus 前端订单支付成功回调-修改订单状态
// alitrip.merchant.galaxy.derby.member.voucher.update.status
//
// 前端订单支付成功回调-修改订单状态
func Alitripmerchantgalaxyderbymembervoucherupdatestatus(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervoucherupdatestatusAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervoucherupdatestatusAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervoucherupdatestatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
