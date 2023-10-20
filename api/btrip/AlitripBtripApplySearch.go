package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripApplySearch 搜索审批单
// alitrip.btrip.apply.search
//
// 外部企业调用获取本企业审批单列表数据
func AlitripBtripApplySearch(clt *core.SDKClient, req *btrip.AlitripBtripApplySearchAPIRequest, session string) (*btrip.AlitripBtripApplySearchAPIResponse, error) {
	var resp btrip.AlitripBtripApplySearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
