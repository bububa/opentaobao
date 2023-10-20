package xhotelcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelcrm"
)

// Taobaoxhotelmemberderbycouponsend 发券
// taobao.xhotel.member.derby.coupon.send
//
// 发券
func Taobaoxhotelmemberderbycouponsend(clt *core.SDKClient, req *xhotelcrm.TaobaoxhotelmemberderbycouponsendAPIRequest, session string) (*xhotelcrm.TaobaoxhotelmemberderbycouponsendAPIResponse, error) {
	var resp xhotelcrm.TaobaoxhotelmemberderbycouponsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
