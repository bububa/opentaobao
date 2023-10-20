package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripProjectAdd 添加项目
// alitrip.btrip.project.add
//
// 添加项目
func AlitripBtripProjectAdd(clt *core.SDKClient, req *btrip.AlitripBtripProjectAddAPIRequest, resp *btrip.AlitripBtripProjectAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
