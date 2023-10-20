package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkseriesskuremove 系列品商品变更-移除商品
// alibaba.wdk.series.sku.remove
//
// 系列品商品变更-移除商品
func Alibabawdkseriesskuremove(clt *core.SDKClient, req *wdk.AlibabawdkseriesskuremoveAPIRequest, session string) (*wdk.AlibabawdkseriesskuremoveAPIResponse, error) {
	var resp wdk.AlibabawdkseriesskuremoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
