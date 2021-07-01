package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBaichuanAsoActivateAPIRequest
设备安装活动激活 API请求
alibaba.baichuan.aso.activate

设备安装活动激活 */
type AlibabaBaichuanAsoActivateAPIRequest struct {
	model.Params
	// 来源
	_source string
	// 1-tmail,2-taobao
	_appId string
	// 1-android,2-ios
	_appOs int64
	// 设备信息，ios为idfa ，android 为imei + imsi
	_deviceInfo *AsoDeviceInfoDo
}

// New
