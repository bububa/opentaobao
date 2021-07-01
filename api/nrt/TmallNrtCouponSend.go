package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

/* TmallNrtCouponSend
券发放接口
tmall.nrt.coupon.send

新零售场景，商家自有渠道发放券 */
func TmallNrtCouponSend(clt *core.SDKClient, req *nrt.TmallNrtCouponSendAPIRequest, session string) (*nrt.TmallNrtCouponSendAPIResponse, error) {
	var resp nrt.TmallNrtCouponSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
