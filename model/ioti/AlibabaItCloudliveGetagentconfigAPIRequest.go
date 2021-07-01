package ioti

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItCloudliveGetagentconfigAPIRequest
线上巡店Agent获取配置 API请求
alibaba.it.cloudlive.getagentconfig

线上巡店应用，外部Agent设备获取设备配置信息，根据配置信息链接mqtt，跟云端进行进一步的消息通信。 */
type AlibabaItCloudliveGetagentconfigAPIRequest struct {
	model.Params
	// agent标识信息
	_agentId string
	// 时间戳
	_timeStamp int64
	// 签名
	_signature string
	// 设备所在IP地址
	_agentIp string
}

// New
