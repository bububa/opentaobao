package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripApplyGet 获取单个审批单
// alitrip.btrip.apply.get
//
// 获取单个审批单的详情数据
func AlitripBtripApplyGet(clt *core.SDKClient, req *btrip.AlitripBtripApplyGetAPIRequest, resp *btrip.AlitripBtripApplyGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
