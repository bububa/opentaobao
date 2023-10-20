package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// Aliexpresstradebuyplaceorder AE下单API
// aliexpress.trade.buy.placeorder
//
// 150欧欧盟税改
func Aliexpresstradebuyplaceorder(clt *core.SDKClient, req *aedropshiper.AliexpresstradebuyplaceorderAPIRequest, session string) (*aedropshiper.AliexpresstradebuyplaceorderAPIResponse, error) {
	var resp aedropshiper.AliexpresstradebuyplaceorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
