package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkchannelordercreate 创建订单
// alibaba.wdk.channel.order.create
//
// 外部商家创建订单
func Alibabawdkchannelordercreate(clt *core.SDKClient, req *wdk.AlibabawdkchannelordercreateAPIRequest, session string) (*wdk.AlibabawdkchannelordercreateAPIResponse, error) {
	var resp wdk.AlibabawdkchannelordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
