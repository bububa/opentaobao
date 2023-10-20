package wms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoBimTradeorderConsignAPIResponse 驱动保税交易订单发货 API返回值
// cainiao.bim.tradeorder.consign
//
// 驱动保税交易订单发货
type CainiaoBimTradeorderConsignAPIResponse struct {
	model.CommonResponse
	CainiaoBimTradeorderConsignAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoBimTradeorderConsignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoBimTradeorderConsignAPIResponseModel).Reset()
}

// CainiaoBimTradeorderConsignAPIResponseModel is 驱动保税交易订单发货 成功返回结果
type CainiaoBimTradeorderConsignAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_bim_tradeorder_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 菜鸟仓库作业单据号
	StoreOrderCode string `json:"store_order_code,omitempty" xml:"store_order_code,omitempty"`
	// 菜鸟物流订单号,格式为LP开头
	LgOrderCode string `json:"lg_order_code,omitempty" xml:"lg_order_code,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoBimTradeorderConsignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StoreOrderCode = ""
	m.LgOrderCode = ""
}

var poolCainiaoBimTradeorderConsignAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoBimTradeorderConsignAPIResponse)
	},
}

// GetCainiaoBimTradeorderConsignAPIResponse 从 sync.Pool 获取 CainiaoBimTradeorderConsignAPIResponse
func GetCainiaoBimTradeorderConsignAPIResponse() *CainiaoBimTradeorderConsignAPIResponse {
	return poolCainiaoBimTradeorderConsignAPIResponse.Get().(*CainiaoBimTradeorderConsignAPIResponse)
}

// ReleaseCainiaoBimTradeorderConsignAPIResponse 将 CainiaoBimTradeorderConsignAPIResponse 保存到 sync.Pool
func ReleaseCainiaoBimTradeorderConsignAPIResponse(v *CainiaoBimTradeorderConsignAPIResponse) {
	v.Reset()
	poolCainiaoBimTradeorderConsignAPIResponse.Put(v)
}
