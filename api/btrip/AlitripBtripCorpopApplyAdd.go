package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopApplyAdd 【商旅】isv添加审批单
// alitrip.btrip.corpop.apply.add
//
// 【商旅】isv添加审批单
func AlitripBtripCorpopApplyAdd(clt *core.SDKClient, req *btrip.AlitripBtripCorpopApplyAddAPIRequest, resp *btrip.AlitripBtripCorpopApplyAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
