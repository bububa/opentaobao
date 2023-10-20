package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaoeticketmerchantmadelay 凭证延期
// taobao.eticket.merchant.ma.delay
//
// 订单延期
func Taobaoeticketmerchantmadelay(clt *core.SDKClient, req *eticket.TaobaoeticketmerchantmadelayAPIRequest, session string) (*eticket.TaobaoeticketmerchantmadelayAPIResponse, error) {
	var resp eticket.TaobaoeticketmerchantmadelayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
