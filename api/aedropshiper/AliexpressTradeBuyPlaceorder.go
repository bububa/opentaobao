package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressTradeBuyPlaceorder AE下单API
// aliexpress.trade.buy.placeorder
//
// 150欧欧盟税改
func AliexpressTradeBuyPlaceorder(clt *core.SDKClient, req *aedropshiper.AliexpressTradeBuyPlaceorderAPIRequest, resp *aedropshiper.AliexpressTradeBuyPlaceorderAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
