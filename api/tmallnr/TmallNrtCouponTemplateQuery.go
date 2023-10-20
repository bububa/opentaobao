package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrtcoupontemplatequery 查询券模版
// tmall.nrt.coupon.template.query
//
// 查询券模版
func Tmallnrtcoupontemplatequery(clt *core.SDKClient, req *tmallnr.TmallnrtcoupontemplatequeryAPIRequest, session string) (*tmallnr.TmallnrtcoupontemplatequeryAPIResponse, error) {
	var resp tmallnr.TmallnrtcoupontemplatequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
