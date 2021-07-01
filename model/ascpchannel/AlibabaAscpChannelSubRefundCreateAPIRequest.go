package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpChannelSubRefundCreateAPIRequest
淘外分销逆向创单（子单退） API请求
alibaba.ascp.channel.sub.refund.create

淘外分销逆向创单（子单退） */
type AlibabaAscpChannelSubRefundCreateAPIRequest struct {
	model.Params
	// 子单退款创建请求
	_subRefundCreateReq *ExternalCreateRefundOrderDetailRequest
}

// New
