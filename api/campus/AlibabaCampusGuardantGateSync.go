package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardantGateSync 网点数据同步
// alibaba.campus.guardant.gate.sync
//
// 门禁供应商创建网点同步
func AlibabaCampusGuardantGateSync(clt *core.SDKClient, req *campus.AlibabaCampusGuardantGateSyncAPIRequest, session string) (*campus.AlibabaCampusGuardantGateSyncAPIResponse, error) {
	var resp campus.AlibabaCampusGuardantGateSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
