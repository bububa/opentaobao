package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopApplySearch 【商旅】搜索审批单列表
// alitrip.btrip.corpop.apply.search
//
// 【商旅】搜索审批单列表
func AlitripBtripCorpopApplySearch(clt *core.SDKClient, req *btrip.AlitripBtripCorpopApplySearchAPIRequest, resp *btrip.AlitripBtripCorpopApplySearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
