package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusguardantgatesync 网点数据同步
// alibaba.campus.guardant.gate.sync
//
// 门禁供应商创建网点同步
func Alibabacampusguardantgatesync(clt *core.SDKClient, req *campus.AlibabacampusguardantgatesyncAPIRequest, session string) (*campus.AlibabacampusguardantgatesyncAPIResponse, error) {
	var resp campus.AlibabacampusguardantgatesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
