package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderFutureInfoPutAPIRequest
订单信息上传更新 API请求
taobao.xhotel.order.future.info.put

商家调用推送信息给飞猪平台。 支持如下操作类型：21: 订单状态更新（商家推送订单状态变更）23：酒店房态信息上传（上传一段时间内的酒店房态）25：在线开发票请求确认 26：自助选房请求进行请求确认   27：自助checkIn请求进行请求确认 32: 扫脸入住入住信息回传 （飞猪将登记至公安系统） */
type TaobaoXhotelOrderFutureInfoPutAPIRequest struct {
	model.Params
	// 商家请求流水号
	_outUuid string
	// 操作类型 21: 订单状态更新（商家推送订单状态变更）23：酒店房态信息上传（上传一段时间内的酒店房态）25：在线开发票请求确认 26：自助选房请求进行请求确认   27：自助checkIn请求进行请求确认 32: 扫脸入住入住信息回传 （飞猪将登记至公安系统）
	_operateType int64
	// 酒店编码
	_hotelCode string
	// 字段详细介绍参见 https://open.alitrip.com/docs/doc.htm?docType=1&articleId=106153
	_context string
	// 商家vendor信息。具体值咨询淘宝技术
	_vendor string
	// 请求流水号
	_requestId string
}

// New
