package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaoeticketmerchantmaavailable 电子凭证核销前校验接口
// taobao.eticket.merchant.ma.available
//
// 商家验码之前的调用接口，用来判断是否可以进行核销操作
func Taobaoeticketmerchantmaavailable(clt *core.SDKClient, req *eticket.TaobaoeticketmerchantmaavailableAPIRequest, session string) (*eticket.TaobaoeticketmerchantmaavailableAPIResponse, error) {
	var resp eticket.TaobaoeticketmerchantmaavailableAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
