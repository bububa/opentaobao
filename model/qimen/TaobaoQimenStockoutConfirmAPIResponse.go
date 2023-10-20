package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStockoutConfirmAPIResponse 出库单确认接口 API返回值
// taobao.qimen.stockout.confirm
//
// 货品出库后，WMS将状态回传给ERP
type TaobaoQimenStockoutConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoQimenStockoutConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenStockoutConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenStockoutConfirmAPIResponseModel).Reset()
}

// TaobaoQimenStockoutConfirmAPIResponseModel is 出库单确认接口 成功返回结果
type TaobaoQimenStockoutConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_stockout_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenStockoutConfirmStruct `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenStockoutConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenStockoutConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStockoutConfirmAPIResponse)
	},
}

// GetTaobaoQimenStockoutConfirmAPIResponse 从 sync.Pool 获取 TaobaoQimenStockoutConfirmAPIResponse
func GetTaobaoQimenStockoutConfirmAPIResponse() *TaobaoQimenStockoutConfirmAPIResponse {
	return poolTaobaoQimenStockoutConfirmAPIResponse.Get().(*TaobaoQimenStockoutConfirmAPIResponse)
}

// ReleaseTaobaoQimenStockoutConfirmAPIResponse 将 TaobaoQimenStockoutConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenStockoutConfirmAPIResponse(v *TaobaoQimenStockoutConfirmAPIResponse) {
	v.Reset()
	poolTaobaoQimenStockoutConfirmAPIResponse.Put(v)
}
