package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripCorpopApplySearch
【商旅】搜索审批单列表
alitrip.btrip.corpop.apply.search

【商旅】搜索审批单列表 */
func AlitripBtripCorpopApplySearch(clt *core.SDKClient, req *btrip.AlitripBtripCorpopApplySearchAPIRequest, session string) (*btrip.AlitripBtripCorpopApplySearchAPIResponse, error) {
	var resp btrip.AlitripBtripCorpopApplySearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
