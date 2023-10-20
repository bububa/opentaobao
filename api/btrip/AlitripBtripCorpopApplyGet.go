package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopApplyGet 【商旅】查询审批单
// alitrip.btrip.corpop.apply.get
//
// 【商旅】查询审批单
func AlitripBtripCorpopApplyGet(clt *core.SDKClient, req *btrip.AlitripBtripCorpopApplyGetAPIRequest, resp *btrip.AlitripBtripCorpopApplyGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
