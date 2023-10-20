package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrtNewcouponSend 券发放接口
// tmall.nrt.newcoupon.send
//
// 券发放接口
func TmallNrtNewcouponSend(clt *core.SDKClient, req *tmallnr.TmallNrtNewcouponSendAPIRequest, session string) (*tmallnr.TmallNrtNewcouponSendAPIResponse, error) {
	var resp tmallnr.TmallNrtNewcouponSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
