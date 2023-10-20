package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtCouponTemplateSyn 喵零券同步
// tmall.nrt.coupon.template.syn
//
// 查询券模版
func TmallNrtCouponTemplateSyn(clt *core.SDKClient, req *nrt.TmallNrtCouponTemplateSynAPIRequest, session string) (*nrt.TmallNrtCouponTemplateSynAPIResponse, error) {
	var resp nrt.TmallNrtCouponTemplateSynAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
