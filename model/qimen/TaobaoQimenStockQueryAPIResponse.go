package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStockQueryAPIResponse 库存查询接口（多条件） API返回值
// taobao.qimen.stock.query
//
// ERP调用奇门的接口,查询商品的库存量
type TaobaoQimenStockQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenStockQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenStockQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenStockQueryAPIResponseModel).Reset()
}

// TaobaoQimenStockQueryAPIResponseModel is 库存查询接口（多条件） 成功返回结果
type TaobaoQimenStockQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_stock_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *StockQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenStockQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenStockQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStockQueryAPIResponse)
	},
}

// GetTaobaoQimenStockQueryAPIResponse 从 sync.Pool 获取 TaobaoQimenStockQueryAPIResponse
func GetTaobaoQimenStockQueryAPIResponse() *TaobaoQimenStockQueryAPIResponse {
	return poolTaobaoQimenStockQueryAPIResponse.Get().(*TaobaoQimenStockQueryAPIResponse)
}

// ReleaseTaobaoQimenStockQueryAPIResponse 将 TaobaoQimenStockQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenStockQueryAPIResponse(v *TaobaoQimenStockQueryAPIResponse) {
	v.Reset()
	poolTaobaoQimenStockQueryAPIResponse.Put(v)
}
