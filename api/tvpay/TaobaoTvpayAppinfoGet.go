package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// Taobaotvpayappinfoget tv支付获取应用信息
// taobao.tvpay.appinfo.get
//
// tv支付获取应用信息
func Taobaotvpayappinfoget(clt *core.SDKClient, req *tvpay.TaobaotvpayappinfogetAPIRequest, session string) (*tvpay.TaobaotvpayappinfogetAPIResponse, error) {
	var resp tvpay.TaobaotvpayappinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
