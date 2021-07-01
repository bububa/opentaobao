package mc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMcRecordOrderSyncAPIRequest
订单信息同步 API请求
tmall.mc.record.order.sync

订单信息同步(零售云接口) */
type TmallMcRecordOrderSyncAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
	// 原价
	_originPrice int64
	// 实付价
	_payPrice int64
	// 用户openId
	_openId string
	// 核销结果
	_result string
	// 云码版本号
	_version string
}

// New
