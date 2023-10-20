package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkstockrealquery 仓内实时库存查询
// alibaba.wdk.stock.real.query
//
// 查询仓内实时库存信息
func Alibabawdkstockrealquery(clt *core.SDKClient, req *wdk.AlibabawdkstockrealqueryAPIRequest, session string) (*wdk.AlibabawdkstockrealqueryAPIResponse, error) {
	var resp wdk.AlibabawdkstockrealqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
