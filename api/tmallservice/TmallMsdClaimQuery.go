package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallmsdclaimquery 查询待理赔工单数据接口
// tmall.msd.claim.query
//
// 查询待理赔工单数据接口
func Tmallmsdclaimquery(clt *core.SDKClient, req *tmallservice.TmallmsdclaimqueryAPIRequest, session string) (*tmallservice.TmallmsdclaimqueryAPIResponse, error) {
	var resp tmallservice.TmallmsdclaimqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
