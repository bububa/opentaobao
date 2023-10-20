package cainiaocntec

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest) Reset() {
	r._queryActivityDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest) GetApiMethodName() string {
	return "cainiao.cntec.shopkeeper.supply.statistics.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryActivityDto is QueryActivityDto Setter
// 查询参数
func (r *CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest) SetQueryActivityDto(_queryActivityDto *QueryActivityDto) error {
	r._queryActivityDto = _queryActivityDto
	r.Set("query_activity_dto", _queryActivityDto)
	return nil
}

// GetQueryActivityDto QueryActivityDto Getter
func (r CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest) GetQueryActivityDto() *QueryActivityDto {
	return r._queryActivityDto
}

var poolCainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoCntecShopkeeperSupplyStatisticsQueryRequest()
	},
}

// GetCainiaoCntecShopkeeperSupplyStatisticsQueryRequest 从 sync.Pool 获取 CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest
func GetCainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest() *CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest {
	return poolCainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest.Get().(*CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest)
}

// ReleaseCainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest 将 CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest 放入 sync.Pool
func ReleaseCainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest(v *CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest) {
	v.Reset()
	poolCainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest.Put(v)
}
