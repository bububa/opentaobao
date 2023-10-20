package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkseriesdefaultskureset 系列品商品变更-重置默认商品
// alibaba.wdk.series.defaultsku.reset
//
// 系列品商品变更-重置默认商品
func Alibabawdkseriesdefaultskureset(clt *core.SDKClient, req *wdk.AlibabawdkseriesdefaultskuresetAPIRequest, session string) (*wdk.AlibabawdkseriesdefaultskuresetAPIResponse, error) {
	var resp wdk.AlibabawdkseriesdefaultskuresetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
