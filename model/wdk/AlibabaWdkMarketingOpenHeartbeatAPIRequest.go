package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingOpenHeartbeatAPIRequest
心跳服务【10s一次】 API请求
alibaba.wdk.marketing.open.heartbeat

商家数据同步心跳服务 */
type AlibabaWdkMarketingOpenHeartbeatAPIRequest struct {
	model.Params
	// 心跳信息
	_heartBeat *HeartBeatBo
}

// New
