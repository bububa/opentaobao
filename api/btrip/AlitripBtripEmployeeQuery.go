package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripemployeequery 企业员工查询
// alitrip.btrip.employee.query
//
// 企业员工查询
func Alitripbtripemployeequery(clt *core.SDKClient, req *btrip.AlitripbtripemployeequeryAPIRequest, session string) (*btrip.AlitripbtripemployeequeryAPIResponse, error) {
	var resp btrip.AlitripbtripemployeequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
