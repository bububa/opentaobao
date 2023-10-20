package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripprojectdelete 删除项目
// alitrip.btrip.project.delete
//
// 删除项目
func Alitripbtripprojectdelete(clt *core.SDKClient, req *btrip.AlitripbtripprojectdeleteAPIRequest, session string) (*btrip.AlitripbtripprojectdeleteAPIResponse, error) {
	var resp btrip.AlitripbtripprojectdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
