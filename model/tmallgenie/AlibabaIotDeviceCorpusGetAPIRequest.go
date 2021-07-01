package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIotDeviceCorpusGetAPIRequest
IoT设备支持语料获取 API请求
alibaba.iot.device.corpus.get

ISV通过该接口获取天猫精灵IoT设备支持控制或查询的语料 */
type AlibabaIotDeviceCorpusGetAPIRequest struct {
	model.Params
	// 天猫精灵开放用户id
	_userOpenId string
	// 天猫精灵开放的client id
	_clientId string
	// iot设备id
	_devId string
}

// New
