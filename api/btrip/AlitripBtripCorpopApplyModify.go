package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopApplyModify 【商旅】修改出差审批单（行程）
// alitrip.btrip.corpop.apply.modify
//
// 【商旅】修改出差审批单（行程）
func AlitripBtripCorpopApplyModify(clt *core.SDKClient, req *btrip.AlitripBtripCorpopApplyModifyAPIRequest, resp *btrip.AlitripBtripCorpopApplyModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
