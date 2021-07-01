package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosOcTradeSyncbanksaleAPIRequest
云闪付、银行卡销售数据回传接口 API请求
alibaba.mos.oc.trade.syncbanksale

云闪付、银行卡销售数据回传 */
type AlibabaMosOcTradeSyncbanksaleAPIRequest struct {
	model.Params
	// pos云闪付、银行卡销售数据
	_posBankSaleInfoDto *PosBankSaleInfoDto
}

// New
