package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripCorpopApplyGet
【商旅】查询审批单
alitrip.btrip.corpop.apply.get

【商旅】查询审批单 */
func AlitripBtripCorpopApplyGet(clt *core.SDKClient, req *btrip.AlitripBtripCorpopApplyGetAPIRequest, session string) (*btrip.AlitripBtripCorpopApplyGetAPIResponse, error) {
	var resp btrip.AlitripBtripCorpopApplyGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
