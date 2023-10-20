package alitripbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripbp"
)

// Alitripbpcouponinfosync 飞猪广告券信息同步接口
// alitrip.bp.couponinfo.sync
//
// 飞猪商业化券信息同步
func Alitripbpcouponinfosync(clt *core.SDKClient, req *alitripbp.AlitripbpcouponinfosyncAPIRequest, session string) (*alitripbp.AlitripbpcouponinfosyncAPIResponse, error) {
	var resp alitripbp.AlitripbpcouponinfosyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
