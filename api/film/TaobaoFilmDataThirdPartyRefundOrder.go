package film

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/film"
)

// TaobaoFilmDataThirdPartyRefundOrder 退票接口
// taobao.film.data.third.party.refund.order
//
// 淘票票第三方退票接口
func TaobaoFilmDataThirdPartyRefundOrder(clt *core.SDKClient, req *film.TaobaoFilmDataThirdPartyRefundOrderAPIRequest, resp *film.TaobaoFilmDataThirdPartyRefundOrderAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
