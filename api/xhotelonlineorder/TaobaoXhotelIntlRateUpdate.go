package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelIntlRateUpdate 不落库商家推送更新酒店rate
// taobao.xhotel.intl.rate.update
//
// 商家主动推送不落库商品的酒店信息
func TaobaoXhotelIntlRateUpdate(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelIntlRateUpdateAPIRequest, resp *xhotelonlineorder.TaobaoXhotelIntlRateUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
