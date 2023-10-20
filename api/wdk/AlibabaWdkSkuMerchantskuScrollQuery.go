package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskumerchantskuscrollquery 商家商品批量查询接口
// alibaba.wdk.sku.merchantsku.scroll.query
//
// 提供主档商品数据接口查询
func Alibabawdkskumerchantskuscrollquery(clt *core.SDKClient, req *wdk.AlibabawdkskumerchantskuscrollqueryAPIRequest, session string) (*wdk.AlibabawdkskumerchantskuscrollqueryAPIResponse, error) {
	var resp wdk.AlibabawdkskumerchantskuscrollqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
