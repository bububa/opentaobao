package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskuquery 查询商品
// alibaba.wdk.sku.query
//
// 查询商品
func Alibabawdkskuquery(clt *core.SDKClient, req *wdk.AlibabawdkskuqueryAPIRequest, session string) (*wdk.AlibabawdkskuqueryAPIResponse, error) {
	var resp wdk.AlibabawdkskuqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
