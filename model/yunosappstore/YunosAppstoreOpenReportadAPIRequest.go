package yunosappstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosAppstoreOpenReportadAPIRequest
外投广告上报接口 API请求
yunos.appstore.open.reportad

外投广告回流上报接口 */
type YunosAppstoreOpenReportadAPIRequest struct {
	model.Params
	// 广告跟踪id列表
	_traceIds []string
	// 事件时间，当前毫秒数
	_time int64
	// 事件类型：0 代表曝光事件；1 代表点击下载事件；2 代表下载完成事件；3 代表安装完成事件
	_action int64
	// 客户端版本号
	_clientVerCode int64
	// 客户端设备标识
	_deviceId string
}

// New
