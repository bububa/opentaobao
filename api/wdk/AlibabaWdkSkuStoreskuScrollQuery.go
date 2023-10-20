package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskustoreskuscrollquery 门店商品批量查询接口
// alibaba.wdk.sku.storesku.scroll.query
//
// 提供门店商品批量查询接口
func Alibabawdkskustoreskuscrollquery(clt *core.SDKClient, req *wdk.AlibabawdkskustoreskuscrollqueryAPIRequest, session string) (*wdk.AlibabawdkskustoreskuscrollqueryAPIResponse, error) {
	var resp wdk.AlibabawdkskustoreskuscrollqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
