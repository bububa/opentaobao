package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopDepartSync 外部部门同步
// alitrip.btrip.corpop.depart.sync
//
// 同步外部平台部门信息至商旅内部
func AlitripBtripCorpopDepartSync(clt *core.SDKClient, req *btrip.AlitripBtripCorpopDepartSyncAPIRequest, resp *btrip.AlitripBtripCorpopDepartSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
