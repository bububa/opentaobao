package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// TaobaoTvpayAccessDataGet tv支付
// taobao.tvpay.access.data.get
//
// 在匿名用户支付后尝试为其登陆绑定的淘宝账号
func TaobaoTvpayAccessDataGet(clt *core.SDKClient, req *tvpay.TaobaoTvpayAccessDataGetAPIRequest, session string) (*tvpay.TaobaoTvpayAccessDataGetAPIResponse, error) {
	var resp tvpay.TaobaoTvpayAccessDataGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
