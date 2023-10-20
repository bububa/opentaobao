package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkstockpublish 五道口库存发布接口（针对外部渠道）
// alibaba.wdk.stock.publish
//
// 五道口库存发布接口（针对外部渠道）
func Alibabawdkstockpublish(clt *core.SDKClient, req *wdk.AlibabawdkstockpublishAPIRequest, session string) (*wdk.AlibabawdkstockpublishAPIResponse, error) {
	var resp wdk.AlibabawdkstockpublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
