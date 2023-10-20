package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// Taobaotvpayaccessdataget tv支付
// taobao.tvpay.access.data.get
//
// 在匿名用户支付后尝试为其登陆绑定的淘宝账号
func Taobaotvpayaccessdataget(clt *core.SDKClient, req *tvpay.TaobaotvpayaccessdatagetAPIRequest, session string) (*tvpay.TaobaotvpayaccessdatagetAPIResponse, error) {
	var resp tvpay.TaobaotvpayaccessdatagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
