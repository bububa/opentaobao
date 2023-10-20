package aliexpresssumaitong

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliexpresssumaitong"
)

// Aliexpresstradeorderopenquery Aliexpress开放平台订单查询
// aliexpress.trade.order.open.query
//
// Aliexpress开放平台订单信息查询
func Aliexpresstradeorderopenquery(clt *core.SDKClient, req *aliexpresssumaitong.AliexpresstradeorderopenqueryAPIRequest, session string) (*aliexpresssumaitong.AliexpresstradeorderopenqueryAPIResponse, error) {
	var resp aliexpresssumaitong.AliexpresstradeorderopenqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
