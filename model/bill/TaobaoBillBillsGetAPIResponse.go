package bill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBillBillsGetAPIResponse 查询账单明细数据(自研发商家专用) API返回值
// taobao.bill.bills.get
//
// 查询账单明细数据
type TaobaoBillBillsGetAPIResponse struct {
	model.CommonResponse
	TaobaoBillBillsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBillBillsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBillBillsGetAPIResponseModel).Reset()
}

// TaobaoBillBillsGetAPIResponseModel is 查询账单明细数据(自研发商家专用) 成功返回结果
type TaobaoBillBillsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"bill_bills_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 账单列表
	Bills []Bill `json:"bills,omitempty" xml:"bills>bill,omitempty"`
	// 当前页查询返回的结果数(0-100)。相同的查询时间段条件下，最大只能获取总共5000条记录。所以当大于等于5000时 ISV可以通过start_time及end_time来进行拆分，以保证可以查询到全部数据
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否存在下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBillBillsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Bills = m.Bills[:0]
	m.TotalResults = 0
	m.HasNext = false
}

var poolTaobaoBillBillsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBillBillsGetAPIResponse)
	},
}

// GetTaobaoBillBillsGetAPIResponse 从 sync.Pool 获取 TaobaoBillBillsGetAPIResponse
func GetTaobaoBillBillsGetAPIResponse() *TaobaoBillBillsGetAPIResponse {
	return poolTaobaoBillBillsGetAPIResponse.Get().(*TaobaoBillBillsGetAPIResponse)
}

// ReleaseTaobaoBillBillsGetAPIResponse 将 TaobaoBillBillsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBillBillsGetAPIResponse(v *TaobaoBillBillsGetAPIResponse) {
	v.Reset()
	poolTaobaoBillBillsGetAPIResponse.Put(v)
}
