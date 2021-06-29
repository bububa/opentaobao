package cainiaocntec

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
团购业务供货商查询门店统计数据 APIRequest
cainiao.cntec.shopkeeper.supply.statistics.query

查询门店售卖商品统计数据
*/
type CainiaoCntecShopkeeperSupplyStatisticsQueryRequest struct {
    model.Params

    // 查询参数
    queryActivityDto   *QueryActivityDto 

}

func NewCainiaoCntecShopkeeperSupplyStatisticsQueryRequest() *CainiaoCntecShopkeeperSupplyStatisticsQueryRequest{
    return &CainiaoCntecShopkeeperSupplyStatisticsQueryRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCntecShopkeeperSupplyStatisticsQueryRequest) GetApiMethodName() string {
    return "cainiao.cntec.shopkeeper.supply.statistics.query"
}

func (r CainiaoCntecShopkeeperSupplyStatisticsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCntecShopkeeperSupplyStatisticsQueryRequest) SetQueryActivityDto(queryActivityDto *QueryActivityDto) error {
    r.queryActivityDto = queryActivityDto
    r.Set("query_activity_dto", queryActivityDto)
    return nil
}

func (r CainiaoCntecShopkeeperSupplyStatisticsQueryRequest) GetQueryActivityDto() *QueryActivityDto {
    return r.queryActivityDto
}

