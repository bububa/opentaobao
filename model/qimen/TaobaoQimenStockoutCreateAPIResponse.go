package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStockoutCreateAPIResponse 出库单创建接口 API返回值
// taobao.qimen.stockout.create
//
// taobao.qimen.returnorder.create
type TaobaoQimenStockoutCreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenStockoutCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenStockoutCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenStockoutCreateAPIResponseModel).Reset()
}

// TaobaoQimenStockoutCreateAPIResponseModel is 出库单创建接口 成功返回结果
type TaobaoQimenStockoutCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_stockout_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenStockoutCreateResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenStockoutCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenStockoutCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStockoutCreateAPIResponse)
	},
}

// GetTaobaoQimenStockoutCreateAPIResponse 从 sync.Pool 获取 TaobaoQimenStockoutCreateAPIResponse
func GetTaobaoQimenStockoutCreateAPIResponse() *TaobaoQimenStockoutCreateAPIResponse {
	return poolTaobaoQimenStockoutCreateAPIResponse.Get().(*TaobaoQimenStockoutCreateAPIResponse)
}

// ReleaseTaobaoQimenStockoutCreateAPIResponse 将 TaobaoQimenStockoutCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenStockoutCreateAPIResponse(v *TaobaoQimenStockoutCreateAPIResponse) {
	v.Reset()
	poolTaobaoQimenStockoutCreateAPIResponse.Put(v)
}
