package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkfulfillbatchontaskstatuschanged 物流管控作业状态回传
// alibaba.wdk.fulfill.batch.on.task.status.changed
//
// 物流管控作业状态回传
func Alibabawdkfulfillbatchontaskstatuschanged(clt *core.SDKClient, req *wdk.AlibabawdkfulfillbatchontaskstatuschangedAPIRequest, session string) (*wdk.AlibabawdkfulfillbatchontaskstatuschangedAPIResponse, error) {
	var resp wdk.AlibabawdkfulfillbatchontaskstatuschangedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
