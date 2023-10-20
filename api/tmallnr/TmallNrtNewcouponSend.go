package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrtNewcouponSend 券发放接口
// tmall.nrt.newcoupon.send
//
// 券发放接口
func TmallNrtNewcouponSend(clt *core.SDKClient, req *tmallnr.TmallNrtNewcouponSendAPIRequest, resp *tmallnr.TmallNrtNewcouponSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
