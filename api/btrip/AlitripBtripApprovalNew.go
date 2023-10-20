package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripapprovalnew 新建审批单
// alitrip.btrip.approval.new
//
// 用户新建审批单
func Alitripbtripapprovalnew(clt *core.SDKClient, req *btrip.AlitripbtripapprovalnewAPIRequest, session string) (*btrip.AlitripbtripapprovalnewAPIResponse, error) {
	var resp btrip.AlitripbtripapprovalnewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
