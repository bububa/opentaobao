package cainiaocntec

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest 团购业务供货商查询门店统计数据 API请求
// cainiao.cntec.shopkeeper.supply.statistics.query
//
// 查询门店售卖商品统计数据
type CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest struct {
	model.Params
	// 查询参数
	_queryActivityDto *QueryActivityDto
}

// NewCainiaoCntecShopkeeperSupplyStatisticsQueryRequest 初始化CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest对象
func NewCainiaoCntecShopkeeperSupplyStatisticsQueryRequest() *CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest {
	return &CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest) GetApiMethodName() string {
	return "cainiao.cntec.shopkeeper.supply.statistics.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is QueryActivityDto Setter
// 查询参数
func (r *CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest) SetQueryActivityDto(_queryActivityDto *QueryActivityDto) error {
	r._queryActivityDto = _queryActivityDto
	r.Set("query_activity_dto", _queryActivityDto)
	return nil
}

// Get QueryActivityDto Getter
func (r CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest) GetQueryActivityDto() *QueryActivityDto {
	return r._queryActivityDto
}
