package cainiaocntec

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocntecshopkeepersupplystatisticsqueryAPIRequest 团购业务供货商查询门店统计数据 API请求
// cainiao.cntec.shopkeeper.supply.statistics.query
//
// 查询门店售卖商品统计数据
type CainiaocntecshopkeepersupplystatisticsqueryAPIRequest struct {
	model.Params
	// 查询参数
	_queryActivityDto *QueryActivityDto
}

// NewCainiaocntecshopkeepersupplystatisticsqueryRequest 初始化CainiaocntecshopkeepersupplystatisticsqueryAPIRequest对象
func NewCainiaocntecshopkeepersupplystatisticsqueryRequest() *CainiaocntecshopkeepersupplystatisticsqueryAPIRequest {
	return &CainiaocntecshopkeepersupplystatisticsqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocntecshopkeepersupplystatisticsqueryAPIRequest) GetApiMethodName() string {
	return "cainiao.cntec.shopkeeper.supply.statistics.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocntecshopkeepersupplystatisticsqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocntecshopkeepersupplystatisticsqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryActivityDto is QueryActivityDto Setter
// 查询参数
func (r *CainiaocntecshopkeepersupplystatisticsqueryAPIRequest) SetQueryActivityDto(_queryActivityDto *QueryActivityDto) error {
	r._queryActivityDto = _queryActivityDto
	r.Set("query_activity_dto", _queryActivityDto)
	return nil
}

// GetQueryActivityDto QueryActivityDto Getter
func (r CainiaocntecshopkeepersupplystatisticsqueryAPIRequest) GetQueryActivityDto() *QueryActivityDto {
	return r._queryActivityDto
}
