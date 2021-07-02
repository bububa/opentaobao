package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// TaobaoTvpayAppinfoGet tv支付获取应用信息
// taobao.tvpay.appinfo.get
//
// tv支付获取应用信息
func TaobaoTvpayAppinfoGet(clt *core.SDKClient, req *tvpay.TaobaoTvpayAppinfoGetAPIRequest, session string) (*tvpay.TaobaoTvpayAppinfoGetAPIResponse, error) {
	var resp tvpay.TaobaoTvpayAppinfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
