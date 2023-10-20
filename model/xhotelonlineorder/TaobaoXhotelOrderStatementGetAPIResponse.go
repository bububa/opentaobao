package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderStatementGetAPIResponse 查询账单信息 API返回值
// taobao.xhotel.order.statement.get
//
// 阿里根据此接口定义输出订单账务明细，结账状态发生变化时阿里需推送账单信息。系统商可实时调用该接口来查询订单的详情
type TaobaoXhotelOrderStatementGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderStatementGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderStatementGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderStatementGetAPIResponseModel).Reset()
}

// TaobaoXhotelOrderStatementGetAPIResponseModel is 查询账单信息 成功返回结果
type TaobaoXhotelOrderStatementGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_statement_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Results []StatementOrder `json:"results,omitempty" xml:"results>statement_order,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderStatementGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoXhotelOrderStatementGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderStatementGetAPIResponse)
	},
}

// GetTaobaoXhotelOrderStatementGetAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderStatementGetAPIResponse
func GetTaobaoXhotelOrderStatementGetAPIResponse() *TaobaoXhotelOrderStatementGetAPIResponse {
	return poolTaobaoXhotelOrderStatementGetAPIResponse.Get().(*TaobaoXhotelOrderStatementGetAPIResponse)
}

// ReleaseTaobaoXhotelOrderStatementGetAPIResponse 将 TaobaoXhotelOrderStatementGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderStatementGetAPIResponse(v *TaobaoXhotelOrderStatementGetAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderStatementGetAPIResponse.Put(v)
}
