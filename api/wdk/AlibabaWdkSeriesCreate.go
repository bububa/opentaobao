package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkseriescreate 系列品变更-新增系列
// alibaba.wdk.series.create
//
// 系列品变更-新增系列
func Alibabawdkseriescreate(clt *core.SDKClient, req *wdk.AlibabawdkseriescreateAPIRequest, session string) (*wdk.AlibabawdkseriescreateAPIResponse, error) {
	var resp wdk.AlibabawdkseriescreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
