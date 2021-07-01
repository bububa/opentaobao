package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripCorpopApplyModify
【商旅】修改出差审批单（行程）
alitrip.btrip.corpop.apply.modify

【商旅】修改出差审批单（行程） */
func AlitripBtripCorpopApplyModify(clt *core.SDKClient, req *btrip.AlitripBtripCorpopApplyModifyAPIRequest, session string) (*btrip.AlitripBtripCorpopApplyModifyAPIResponse, error) {
	var resp btrip.AlitripBtripCorpopApplyModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
