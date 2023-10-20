package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaoeticketmerchantmaconsume 电子凭证核销接口
// taobao.eticket.merchant.ma.consume
//
// 电子凭证核销接口
func Taobaoeticketmerchantmaconsume(clt *core.SDKClient, req *eticket.TaobaoeticketmerchantmaconsumeAPIRequest, session string) (*eticket.TaobaoeticketmerchantmaconsumeAPIResponse, error) {
	var resp eticket.TaobaoeticketmerchantmaconsumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
