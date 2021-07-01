package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBaichuanAsoQueryAPIRequest
查询app在设备上的安装信息 API请求
alibaba.baichuan.aso.query

查询app在设备上的安装信息 */
type AlibabaBaichuanAsoQueryAPIRequest struct {
	model.Params
	// 1-tmail,2-taobao
	_appId string
	// 1-android,2-ios
	_appOs int64
	// 设备信息，ios为idfa ，android 为imei + imsi
	_deviceInfoList []AsoDeviceInfoDo
}

// New
