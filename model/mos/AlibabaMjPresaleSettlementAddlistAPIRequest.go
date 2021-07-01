package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjPresaleSettlementAddlistAPIRequest
预售结算数据回传 API请求
alibaba.mj.presale.settlement.addlist

用于预售活动结算数据的回传。 */
type AlibabaMjPresaleSettlementAddlistAPIRequest struct {
	model.Params
	// 订单json格式数据
	_preSaleRefundJson string
}

// New
