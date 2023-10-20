package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketOplogsGetAPIResponse 电子凭证操作日志查询 API返回值
// taobao.vmarket.eticket.oplogs.get
//
// 电子凭证核销日志查询
type TaobaoVmarketEticketOplogsGetAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketOplogsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketOplogsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVmarketEticketOplogsGetAPIResponseModel).Reset()
}

// TaobaoVmarketEticketOplogsGetAPIResponseModel is 电子凭证操作日志查询 成功返回结果
type TaobaoVmarketEticketOplogsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_oplogs_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作日志列表
	EticketOpLogs []EticketOpLog `json:"eticket_op_logs,omitempty" xml:"eticket_op_logs>eticket_op_log,omitempty"`
	// 符合条件的记录总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketOplogsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EticketOpLogs = m.EticketOpLogs[:0]
	m.TotalResults = 0
}

var poolTaobaoVmarketEticketOplogsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVmarketEticketOplogsGetAPIResponse)
	},
}

// GetTaobaoVmarketEticketOplogsGetAPIResponse 从 sync.Pool 获取 TaobaoVmarketEticketOplogsGetAPIResponse
func GetTaobaoVmarketEticketOplogsGetAPIResponse() *TaobaoVmarketEticketOplogsGetAPIResponse {
	return poolTaobaoVmarketEticketOplogsGetAPIResponse.Get().(*TaobaoVmarketEticketOplogsGetAPIResponse)
}

// ReleaseTaobaoVmarketEticketOplogsGetAPIResponse 将 TaobaoVmarketEticketOplogsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVmarketEticketOplogsGetAPIResponse(v *TaobaoVmarketEticketOplogsGetAPIResponse) {
	v.Reset()
	poolTaobaoVmarketEticketOplogsGetAPIResponse.Put(v)
}
