package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtcoupontemplatequery 券模板查询
// tmall.nrt.coupontemplate.query
//
// 新零售场景，商家拉取在新零售工作台设置的券数据
func Tmallnrtcoupontemplatequery(clt *core.SDKClient, req *nrt.TmallnrtcoupontemplatequeryAPIRequest, session string) (*nrt.TmallnrtcoupontemplatequeryAPIResponse, error) {
	var resp nrt.TmallnrtcoupontemplatequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
