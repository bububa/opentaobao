package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest
预售商家仓出库 API请求
alibaba.ascp.uop.taobao.presalesorder.consignconfirm

预售商家仓出库 */
type AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest struct {
	model.Params
	// 预售订单商家仓出库对象
	_presalesOrderConsignConfirmRequest *Presalesorderconsignconfirmrequest
}

// New
