package dengta

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dengta"
)

// AlibabaPicturesDengtaWxaccountPriceChange 微信公众号价格变化通知
// alibaba.pictures.dengta.wxaccount.price.change
//
// 微信公众号推广价格变更通知接口
func AlibabaPicturesDengtaWxaccountPriceChange(clt *core.SDKClient, req *dengta.AlibabaPicturesDengtaWxaccountPriceChangeAPIRequest, resp *dengta.AlibabaPicturesDengtaWxaccountPriceChangeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
