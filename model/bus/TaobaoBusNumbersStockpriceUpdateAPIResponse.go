package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusNumbersStockpriceUpdateAPIResponse 汽车票更新价格库存 API返回值
// taobao.bus.numbers.stockprice.update
//
// 用于汽车票代理商更新价格库存
type TaobaoBusNumbersStockpriceUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoBusNumbersStockpriceUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusNumbersStockpriceUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusNumbersStockpriceUpdateAPIResponseModel).Reset()
}

// TaobaoBusNumbersStockpriceUpdateAPIResponseModel is 汽车票更新价格库存 成功返回结果
type TaobaoBusNumbersStockpriceUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_numbers_stockprice_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 成功数量
	SuccCount int64 `json:"succ_count,omitempty" xml:"succ_count,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusNumbersStockpriceUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.SuccCount = 0
	m.IsSuccess = false
}

var poolTaobaoBusNumbersStockpriceUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusNumbersStockpriceUpdateAPIResponse)
	},
}

// GetTaobaoBusNumbersStockpriceUpdateAPIResponse 从 sync.Pool 获取 TaobaoBusNumbersStockpriceUpdateAPIResponse
func GetTaobaoBusNumbersStockpriceUpdateAPIResponse() *TaobaoBusNumbersStockpriceUpdateAPIResponse {
	return poolTaobaoBusNumbersStockpriceUpdateAPIResponse.Get().(*TaobaoBusNumbersStockpriceUpdateAPIResponse)
}

// ReleaseTaobaoBusNumbersStockpriceUpdateAPIResponse 将 TaobaoBusNumbersStockpriceUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusNumbersStockpriceUpdateAPIResponse(v *TaobaoBusNumbersStockpriceUpdateAPIResponse) {
	v.Reset()
	poolTaobaoBusNumbersStockpriceUpdateAPIResponse.Put(v)
}
