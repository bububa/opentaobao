package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskuwarehouseskuquery 仓商品查询接口(指定商品编码)
// alibaba.wdk.sku.warehousesku.query
//
// 提供指定仓商品编码查询
func Alibabawdkskuwarehouseskuquery(clt *core.SDKClient, req *wdk.AlibabawdkskuwarehouseskuqueryAPIRequest, session string) (*wdk.AlibabawdkskuwarehouseskuqueryAPIResponse, error) {
	var resp wdk.AlibabawdkskuwarehouseskuqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
