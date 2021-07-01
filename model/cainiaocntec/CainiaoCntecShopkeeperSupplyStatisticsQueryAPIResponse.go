package cainiaocntec

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
团购业务供货商查询门店统计数据 API返回值 
cainiao.cntec.shopkeeper.supply.statistics.query

查询门店售卖商品统计数据
*/
type CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse struct {
    model.CommonResponse
    CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponseModel
}

// 团购业务供货商查询门店统计数据 成功返回结果
type CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"cainiao_cntec_shopkeeper_supply_statistics_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *CainiaoCntecShopkeeperSupplyStatisticsQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
