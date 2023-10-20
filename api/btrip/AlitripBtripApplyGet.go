package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripapplyget 获取单个审批单
// alitrip.btrip.apply.get
//
// 获取单个审批单的详情数据
func Alitripbtripapplyget(clt *core.SDKClient, req *btrip.AlitripbtripapplygetAPIRequest, session string) (*btrip.AlitripbtripapplygetAPIResponse, error) {
	var resp btrip.AlitripbtripapplygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
