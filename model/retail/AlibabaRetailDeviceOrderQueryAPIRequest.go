package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailDeviceOrderQueryAPIRequest
贩卖机订单查询 API请求
alibaba.retail.device.order.query

贩卖机订单查询 */
type AlibabaRetailDeviceOrderQueryAPIRequest struct {
	model.Params
	// 阿里设备物理ID
	_deviceSnList []string
	// 外部设备编码
	_deviceUuid string
	// 阿里设备编码
	_deviceCode string
	// -20 已退款，-10 交易关闭 ，10 创单 20 已支付  30 已出货  40 交易完成
	_status int64
	// CASH 现金，ALIPAY_FACE_PAY_QR 支付宝，VENDING_PRIZE 抽奖，FACE_PAY 人脸
	_payType string
	// 分页大小
	_pageSize int64
	// 页码
	_pageNum int64
	// 查询订单开始时间
	_starts string
	// 查询订单结束时间
	_ends string
}

// New
