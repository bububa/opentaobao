package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripProjectDelete 删除项目
// alitrip.btrip.project.delete
//
// 删除项目
func AlitripBtripProjectDelete(clt *core.SDKClient, req *btrip.AlitripBtripProjectDeleteAPIRequest, resp *btrip.AlitripBtripProjectDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
