package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripApplySearch 搜索审批单
// alitrip.btrip.apply.search
//
// 外部企业调用获取本企业审批单列表数据
func AlitripBtripApplySearch(clt *core.SDKClient, req *btrip.AlitripBtripApplySearchAPIRequest, resp *btrip.AlitripBtripApplySearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
