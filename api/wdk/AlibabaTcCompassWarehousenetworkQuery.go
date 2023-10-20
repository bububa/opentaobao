package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatccompasswarehousenetworkquery 按仓维度来查询鸟潮网络
// alibaba.tc.compass.warehousenetwork.query
//
// 按仓维度来查询鸟潮网络
func Alibabatccompasswarehousenetworkquery(clt *core.SDKClient, req *wdk.AlibabatccompasswarehousenetworkqueryAPIRequest, session string) (*wdk.AlibabatccompasswarehousenetworkqueryAPIResponse, error) {
	var resp wdk.AlibabatccompasswarehousenetworkqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
