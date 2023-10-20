package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallMsdClaimQuery 查询待理赔工单数据接口
// tmall.msd.claim.query
//
// 查询待理赔工单数据接口
func TmallMsdClaimQuery(clt *core.SDKClient, req *tmallservice.TmallMsdClaimQueryAPIRequest, session string) (*tmallservice.TmallMsdClaimQueryAPIResponse, error) {
	var resp tmallservice.TmallMsdClaimQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
