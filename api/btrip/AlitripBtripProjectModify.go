package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripProjectModify 变更项目
// alitrip.btrip.project.modify
//
// 变更项目
func AlitripBtripProjectModify(clt *core.SDKClient, req *btrip.AlitripBtripProjectModifyAPIRequest, resp *btrip.AlitripBtripProjectModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
