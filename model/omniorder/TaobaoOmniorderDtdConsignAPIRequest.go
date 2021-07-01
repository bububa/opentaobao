package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderDtdConsignAPIRequest
门店自送发货 API请求
taobao.omniorder.dtd.consign

该接口触发门店自送发货，推进淘系订单状态为发货，为消费者发送核销码短信，并将物流信息写入订单 */
type TaobaoOmniorderDtdConsignAPIRequest struct {
	model.Params
	// 淘宝订单主订单号
	_mainOrderId int64
	// 发货对应的商户中心门店ID
	_storeId int64
}

// New
