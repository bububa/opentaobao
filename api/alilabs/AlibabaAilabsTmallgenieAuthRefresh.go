package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTmallgenieAuthRefresh 刷新token
// alibaba.ailabs.tmallgenie.auth.refresh
//
// 通过此接口刷新天猫精灵授权token
func AlibabaAilabsTmallgenieAuthRefresh(clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthRefreshAPIRequest, session string) (*alilabs.AlibabaAilabsTmallgenieAuthRefreshAPIResponse, error) {
	var resp alilabs.AlibabaAilabsTmallgenieAuthRefreshAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
