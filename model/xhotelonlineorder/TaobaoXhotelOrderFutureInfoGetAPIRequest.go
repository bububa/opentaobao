package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderFutureInfoGetAPIRequest
获取(查询)订单变更信息 API请求
taobao.xhotel.order.future.info.get

支持操作类型 1.在线开发票请求 3.在线选房请求 4.自助checkIn请求 13.扫脸入住身份信息请求 10.房态信息查询请求 103.通用任务取消指令 */
type TaobaoXhotelOrderFutureInfoGetAPIRequest struct {
	model.Params
	// 请求流水号
	_outUuid string
	// 指定淘宝订单ID。以英文分号隔开的字符串“123455666;123455666;123455666”
	_tids string
	// 酒店编码
	_hotelCode string
	// 系统商分配的身份识别
	_vendor string
	// 操作类型 1.在线开发票请求 3.在线选房请求 4.自助checkIn请求 13.扫脸入住身份信息请求 10.房态信息查询请求 103.通用任务取消指令
	_operateType int64
	// 开始时间
	_createdStart string
	// 结束时间
	_createdEnd string
}

// New
