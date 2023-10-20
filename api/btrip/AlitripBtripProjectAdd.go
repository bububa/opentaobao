package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripprojectadd 添加项目
// alitrip.btrip.project.add
//
// 添加项目
func Alitripbtripprojectadd(clt *core.SDKClient, req *btrip.AlitripbtripprojectaddAPIRequest, session string) (*btrip.AlitripbtripprojectaddAPIResponse, error) {
	var resp btrip.AlitripbtripprojectaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
