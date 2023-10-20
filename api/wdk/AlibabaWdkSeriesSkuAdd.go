package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkseriesskuadd 系列品商品变更-添加商品
// alibaba.wdk.series.sku.add
//
// 系列品商品变更-添加商品
func Alibabawdkseriesskuadd(clt *core.SDKClient, req *wdk.AlibabawdkseriesskuaddAPIRequest, session string) (*wdk.AlibabawdkseriesskuaddAPIResponse, error) {
	var resp wdk.AlibabawdkseriesskuaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
