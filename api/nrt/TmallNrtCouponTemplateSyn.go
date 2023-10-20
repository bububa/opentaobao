package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtcoupontemplatesyn 喵零券同步
// tmall.nrt.coupon.template.syn
//
// 查询券模版
func Tmallnrtcoupontemplatesyn(clt *core.SDKClient, req *nrt.TmallnrtcoupontemplatesynAPIRequest, session string) (*nrt.TmallnrtcoupontemplatesynAPIResponse, error) {
	var resp nrt.TmallnrtcoupontemplatesynAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
