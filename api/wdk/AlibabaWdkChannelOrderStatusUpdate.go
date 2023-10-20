package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkchannelorderstatusupdate 订单状态变更
// alibaba.wdk.channel.order.status.update
//
// 订单状态变更
func Alibabawdkchannelorderstatusupdate(clt *core.SDKClient, req *wdk.AlibabawdkchannelorderstatusupdateAPIRequest, session string) (*wdk.AlibabawdkchannelorderstatusupdateAPIResponse, error) {
	var resp wdk.AlibabawdkchannelorderstatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
