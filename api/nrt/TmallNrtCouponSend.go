package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtcouponsend 券发放接口
// tmall.nrt.coupon.send
//
// 新零售场景，商家自有渠道发放券
func Tmallnrtcouponsend(clt *core.SDKClient, req *nrt.TmallnrtcouponsendAPIRequest, session string) (*nrt.TmallnrtcouponsendAPIResponse, error) {
	var resp nrt.TmallnrtcouponsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
