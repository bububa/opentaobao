package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskuwarehouseskuscrollquery 仓商品遍历接口(游标)
// alibaba.wdk.sku.warehousesku.scroll.query
//
// 提供仓商品数据接口查询
func Alibabawdkskuwarehouseskuscrollquery(clt *core.SDKClient, req *wdk.AlibabawdkskuwarehouseskuscrollqueryAPIRequest, session string) (*wdk.AlibabawdkskuwarehouseskuscrollqueryAPIResponse, error) {
	var resp wdk.AlibabawdkskuwarehouseskuscrollqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
