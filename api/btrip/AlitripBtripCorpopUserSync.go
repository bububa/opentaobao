package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopUserSync 外部人员同步
// alitrip.btrip.corpop.user.sync
//
// 同步外部平台用户信息至商旅内部
func AlitripBtripCorpopUserSync(clt *core.SDKClient, req *btrip.AlitripBtripCorpopUserSyncAPIRequest, resp *btrip.AlitripBtripCorpopUserSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
