package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabUserOpenUidGet access token 获取精灵用户 id
// alibaba.ailab.user.open.uid.get
//
// access token 获取精灵用户 id
func AlibabaAilabUserOpenUidGet(clt *core.SDKClient, req *alilabs.AlibabaAilabUserOpenUidGetAPIRequest, resp *alilabs.AlibabaAilabUserOpenUidGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
