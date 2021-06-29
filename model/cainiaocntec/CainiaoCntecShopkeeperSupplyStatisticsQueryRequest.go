package cainiaocntec

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
团购业务供货商查询门店统计数据 API请求
cainiao.cntec.shopkeeper.supply.statistics.query

查询门店售卖商品统计数据
*/
type CainiaoCntecShopkeeperSupplyStatisticsQueryRequest struct {
    model.Params
    // 查询参数
    _queryActivityDto   *QueryActivityDTO
}

// 初始化CainiaoCntecShopkeeperSupplyStatisticsQueryRequest对象
func NewCainiaoCntecShopkeeperSupplyStatisticsQueryRequest() *CainiaoCntecShopkeeperSupplyStatisticsQueryRequest{
    return &CainiaoCntecShopkeeperSupplyStatisticsQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCntecShopkeeperSupplyStatisticsQueryRequest) GetApiMethodName() string {
    return "cainiao.cntec.shopkeeper.supply.statistics.query"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCntecShopkeeperSupplyStatisticsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryActivityDto Setter
// 查询参数
func (r *CainiaoCntecShopkeeperSupplyStatisticsQueryRequest) SetQueryActivityDto(_queryActivityDto *QueryActivityDTO) error {
    r._queryActivityDto = _queryActivityDto
    r.Set("query_activity_dto", _queryActivityDto)
    return nil
}

// QueryActivityDto Getter
func (r CainiaoCntecShopkeeperSupplyStatisticsQueryRequest) GetQueryActivityDto() *QueryActivityDTO {
    return r._queryActivityDto
}
