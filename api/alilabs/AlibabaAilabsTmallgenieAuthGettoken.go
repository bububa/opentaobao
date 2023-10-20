package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTmallgenieAuthGettoken 设备授权
// alibaba.ailabs.tmallgenie.auth.gettoken
//
// 获取设备授权码
func AlibabaAilabsTmallgenieAuthGettoken(clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthGettokenAPIRequest, resp *alilabs.AlibabaAilabsTmallgenieAuthGettokenAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
