package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingopenheartbeat 心跳服务【10s一次】
// alibaba.wdk.marketing.open.heartbeat
//
// 商家数据同步心跳服务
func Alibabawdkmarketingopenheartbeat(clt *core.SDKClient, req *wdk.AlibabawdkmarketingopenheartbeatAPIRequest, session string) (*wdk.AlibabawdkmarketingopenheartbeatAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingopenheartbeatAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
