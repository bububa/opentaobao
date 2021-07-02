package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTmallgenieAuthSwitchuser 切换用户
// alibaba.ailabs.tmallgenie.auth.switchuser
//
// 设备切换授权用户
func AlibabaAilabsTmallgenieAuthSwitchuser(clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest, session string) (*alilabs.AlibabaAilabsTmallgenieAuthSwitchuserAPIResponse, error) {
	var resp alilabs.AlibabaAilabsTmallgenieAuthSwitchuserAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
