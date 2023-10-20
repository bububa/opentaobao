package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTmallgenieAuthSwitchuser 切换用户
// alibaba.ailabs.tmallgenie.auth.switchuser
//
// 设备切换授权用户
func AlibabaAilabsTmallgenieAuthSwitchuser(clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest, resp *alilabs.AlibabaAilabsTmallgenieAuthSwitchuserAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
