package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskuadd 新增商品
// alibaba.wdk.sku.add
//
// 创建RT门店商品或DC商品
func Alibabawdkskuadd(clt *core.SDKClient, req *wdk.AlibabawdkskuaddAPIRequest, session string) (*wdk.AlibabawdkskuaddAPIResponse, error) {
	var resp wdk.AlibabawdkskuaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
