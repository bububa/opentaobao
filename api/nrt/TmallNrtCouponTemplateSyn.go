package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtCouponTemplateSyn 喵零券同步
// tmall.nrt.coupon.template.syn
//
// 查询券模版
func TmallNrtCouponTemplateSyn(clt *core.SDKClient, req *nrt.TmallNrtCouponTemplateSynAPIRequest, resp *nrt.TmallNrtCouponTemplateSynAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
