package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjPresaleSettlementStatisticsAPIRequest
预购结算数据统计 API请求
alibaba.mj.presale.settlement.statistics

预购结算数据统计 */
type AlibabaMjPresaleSettlementStatisticsAPIRequest struct {
	model.Params
	// 活动期号
	_actionNo int64
	// 外部门店编码
	_storeNo string
}

// New
