package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskucombineskuadd 组合商品新增接口
// alibaba.wdk.sku.combinesku.add
//
// 组合商品新增接口
func Alibabawdkskucombineskuadd(clt *core.SDKClient, req *wdk.AlibabawdkskucombineskuaddAPIRequest, session string) (*wdk.AlibabawdkskucombineskuaddAPIResponse, error) {
	var resp wdk.AlibabawdkskucombineskuaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
