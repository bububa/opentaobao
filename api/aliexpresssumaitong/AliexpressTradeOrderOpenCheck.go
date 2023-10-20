package aliexpresssumaitong

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliexpresssumaitong"
)

// Aliexpresstradeorderopencheck Aliexpress开放平台下单前置检查
// aliexpress.trade.order.open.check
//
// Aliexpress开放平台下单前通过下单入参获取token
func Aliexpresstradeorderopencheck(clt *core.SDKClient, req *aliexpresssumaitong.AliexpresstradeorderopencheckAPIRequest, session string) (*aliexpresssumaitong.AliexpresstradeorderopencheckAPIResponse, error) {
	var resp aliexpresssumaitong.AliexpresstradeorderopencheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
