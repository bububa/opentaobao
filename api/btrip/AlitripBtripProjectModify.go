package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripprojectmodify 变更项目
// alitrip.btrip.project.modify
//
// 变更项目
func Alitripbtripprojectmodify(clt *core.SDKClient, req *btrip.AlitripbtripprojectmodifyAPIRequest, session string) (*btrip.AlitripbtripprojectmodifyAPIResponse, error) {
	var resp btrip.AlitripbtripprojectmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
