package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoFulfillmentAuthCheck 商家鉴权
// tmall.aliauto.fulfillment.auth.check
//
// 商家鉴权
func TmallAliautoFulfillmentAuthCheck(clt *core.SDKClient, req *tmallcar.TmallAliautoFulfillmentAuthCheckAPIRequest, session string) (*tmallcar.TmallAliautoFulfillmentAuthCheckAPIResponse, error) {
	var resp tmallcar.TmallAliautoFulfillmentAuthCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
