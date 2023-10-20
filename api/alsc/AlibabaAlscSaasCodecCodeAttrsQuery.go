package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscSaasCodecCodeAttrsQuery 码业务属性查询
// alibaba.alsc.saas.codec.code.attrs.query
//
// 码业务属性查询
func AlibabaAlscSaasCodecCodeAttrsQuery(clt *core.SDKClient, req *alsc.AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest, resp *alsc.AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
