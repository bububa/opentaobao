package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaoeticketmerchantmareverse 电子凭证冲正接口
// taobao.eticket.merchant.ma.reverse
//
// 电子凭证平台冲正接口
func Taobaoeticketmerchantmareverse(clt *core.SDKClient, req *eticket.TaobaoeticketmerchantmareverseAPIRequest, session string) (*eticket.TaobaoeticketmerchantmareverseAPIResponse, error) {
	var resp eticket.TaobaoeticketmerchantmareverseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
