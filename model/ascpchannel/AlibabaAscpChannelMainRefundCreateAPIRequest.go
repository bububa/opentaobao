package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpChannelMainRefundCreateAPIRequest
淘外分销逆向创单（未发货整单退） API请求
alibaba.ascp.channel.main.refund.create

淘外分销解决方案--订单--逆向创单（未发货整单退） */
type AlibabaAscpChannelMainRefundCreateAPIRequest struct {
	model.Params
	// 逆向单创建请求
	_refundCreateRequest *ExternalCreateRefundOrderRequest
}

// New
