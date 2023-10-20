package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskucombineskuquery 组合商品查询接口
// alibaba.wdk.sku.combinesku.query
//
// 查询组合商品接口
func Alibabawdkskucombineskuquery(clt *core.SDKClient, req *wdk.AlibabawdkskucombineskuqueryAPIRequest, session string) (*wdk.AlibabawdkskucombineskuqueryAPIResponse, error) {
	var resp wdk.AlibabawdkskucombineskuqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
