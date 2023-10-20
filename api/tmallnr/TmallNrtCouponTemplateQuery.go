package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrtCouponTemplateQuery 查询券模版
// tmall.nrt.coupon.template.query
//
// 查询券模版
func TmallNrtCouponTemplateQuery(clt *core.SDKClient, req *tmallnr.TmallNrtCouponTemplateQueryAPIRequest, resp *tmallnr.TmallNrtCouponTemplateQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
