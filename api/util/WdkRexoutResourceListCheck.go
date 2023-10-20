package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Wdkrexoutresourcelistcheck ReX应用中心资源更新检测-外部
// wdk.rexout.resource.list.check
//
// ReX应用中心资源更新检测-外部
func Wdkrexoutresourcelistcheck(clt *core.SDKClient, req *util.WdkrexoutresourcelistcheckAPIRequest, session string) (*util.WdkrexoutresourcelistcheckAPIResponse, error) {
	var resp util.WdkrexoutresourcelistcheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
