package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTcCompassWarehousenetworkQuery 按仓维度来查询鸟潮网络
// alibaba.tc.compass.warehousenetwork.query
//
// 按仓维度来查询鸟潮网络
func AlibabaTcCompassWarehousenetworkQuery(clt *core.SDKClient, req *wdk.AlibabaTcCompassWarehousenetworkQueryAPIRequest, session string) (*wdk.AlibabaTcCompassWarehousenetworkQueryAPIResponse, error) {
	var resp wdk.AlibabaTcCompassWarehousenetworkQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
