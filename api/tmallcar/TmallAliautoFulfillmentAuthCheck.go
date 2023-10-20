package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoFulfillmentAuthCheck 商家鉴权
// tmall.aliauto.fulfillment.auth.check
//
// 商家鉴权
func TmallAliautoFulfillmentAuthCheck(clt *core.SDKClient, req *tmallcar.TmallAliautoFulfillmentAuthCheckAPIRequest, resp *tmallcar.TmallAliautoFulfillmentAuthCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
