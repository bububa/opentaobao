package cainiaocntec

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse 团购业务供货商查询门店统计数据 API返回值
// cainiao.cntec.shopkeeper.supply.statistics.query
//
// 查询门店售卖商品统计数据
type CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse struct {
	model.CommonResponse
	CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponseModel).Reset()
}

// CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponseModel is 团购业务供货商查询门店统计数据 成功返回结果
type CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cntec_shopkeeper_supply_statistics_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CainiaoCntecShopkeeperSupplyStatisticsQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse)
	},
}

// GetCainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse 从 sync.Pool 获取 CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse
func GetCainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse() *CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse {
	return poolCainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse.Get().(*CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse)
}

// ReleaseCainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse 将 CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse(v *CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse) {
	v.Reset()
	poolCainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse.Put(v)
}
