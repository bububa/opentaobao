package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkordersync 五道口外部订单同步
// alibaba.wdk.order.sync
//
// 外部商户使用自助POS下单订单同步到五道口
func Alibabawdkordersync(clt *core.SDKClient, req *wdk.AlibabawdkordersyncAPIRequest, session string) (*wdk.AlibabawdkordersyncAPIResponse, error) {
	var resp wdk.AlibabawdkordersyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
