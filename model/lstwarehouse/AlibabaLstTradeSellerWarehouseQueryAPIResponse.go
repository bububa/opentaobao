package lstwarehouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeSellerWarehouseQueryAPIResponse 供应商-本云商家-仓库查询接口 API返回值
// alibaba.lst.trade.seller.warehouse.query
//
// 查询本地云仓商家的仓库
type AlibabaLstTradeSellerWarehouseQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLstTradeSellerWarehouseQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstTradeSellerWarehouseQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstTradeSellerWarehouseQueryAPIResponseModel).Reset()
}

// AlibabaLstTradeSellerWarehouseQueryAPIResponseModel is 供应商-本云商家-仓库查询接口 成功返回结果
type AlibabaLstTradeSellerWarehouseQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_trade_seller_warehouse_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaLstTradeSellerWarehouseQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstTradeSellerWarehouseQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstTradeSellerWarehouseQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeSellerWarehouseQueryAPIResponse)
	},
}

// GetAlibabaLstTradeSellerWarehouseQueryAPIResponse 从 sync.Pool 获取 AlibabaLstTradeSellerWarehouseQueryAPIResponse
func GetAlibabaLstTradeSellerWarehouseQueryAPIResponse() *AlibabaLstTradeSellerWarehouseQueryAPIResponse {
	return poolAlibabaLstTradeSellerWarehouseQueryAPIResponse.Get().(*AlibabaLstTradeSellerWarehouseQueryAPIResponse)
}

// ReleaseAlibabaLstTradeSellerWarehouseQueryAPIResponse 将 AlibabaLstTradeSellerWarehouseQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstTradeSellerWarehouseQueryAPIResponse(v *AlibabaLstTradeSellerWarehouseQueryAPIResponse) {
	v.Reset()
	poolAlibabaLstTradeSellerWarehouseQueryAPIResponse.Put(v)
}
