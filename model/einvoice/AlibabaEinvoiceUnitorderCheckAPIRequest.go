package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceUnitorderCheckAPIRequest
服务商订购单上传核对 API请求
alibaba.einvoice.unitorder.check

开票服务商回传收到的订购单用于电子发票平台核对 */
type AlibabaEinvoiceUnitorderCheckAPIRequest struct {
	model.Params
	// 订购单列表
	_orders []SimpleUnitOrder
	// 开始时间,来自于查询消息
	_begin string
	// 结束时间,来自于查询消息
	_end string
}

// New
