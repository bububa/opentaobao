package cainiaocntec

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest
团购业务供货商查询门店统计数据 API请求
cainiao.cntec.shopkeeper.supply.statistics.query

查询门店售卖商品统计数据 */
type CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest struct {
	model.Params
	// 查询参数
	_queryActivityDto *QueryActivityDto
}

// New
