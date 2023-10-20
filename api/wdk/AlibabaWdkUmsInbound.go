package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkumsinbound 入库-ERP下发单
// alibaba.wdk.ums.inbound
//
// 入库-ERP下发单
func Alibabawdkumsinbound(clt *core.SDKClient, req *wdk.AlibabawdkumsinboundAPIRequest, session string) (*wdk.AlibabawdkumsinboundAPIResponse, error) {
	var resp wdk.AlibabawdkumsinboundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
