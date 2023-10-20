package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautofulfillmentauthcheck 商家鉴权
// tmall.aliauto.fulfillment.auth.check
//
// 商家鉴权
func Tmallaliautofulfillmentauthcheck(clt *core.SDKClient, req *tmallcar.TmallaliautofulfillmentauthcheckAPIRequest, session string) (*tmallcar.TmallaliautofulfillmentauthcheckAPIResponse, error) {
	var resp tmallcar.TmallaliautofulfillmentauthcheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
