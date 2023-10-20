package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseprojectactivitysync 电商券数据同步
// alibaba.alihouse.project.activity.sync
//
// 电商券数据同步
func Alibabaalihouseprojectactivitysync(clt *core.SDKClient, req *alihouse.AlibabaalihouseprojectactivitysyncAPIRequest, session string) (*alihouse.AlibabaalihouseprojectactivitysyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseprojectactivitysyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
