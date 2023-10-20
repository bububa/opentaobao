package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabasscsupplyplatformserviceworkercancelleave 工人取消请假
// alibaba.ssc.supplyplatform.serviceworker.cancelleave
//
// 工人取消请假
func Alibabasscsupplyplatformserviceworkercancelleave(clt *core.SDKClient, req *tmallservice.AlibabasscsupplyplatformserviceworkercancelleaveAPIRequest, session string) (*tmallservice.AlibabasscsupplyplatformserviceworkercancelleaveAPIResponse, error) {
	var resp tmallservice.AlibabasscsupplyplatformserviceworkercancelleaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
