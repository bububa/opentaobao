package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarFinanceStatusSyncAPIResponse 汽车金融状态同步 API返回值
// tmall.car.finance.status.sync
//
// 汽车金融状态同步
type TmallCarFinanceStatusSyncAPIResponse struct {
	model.CommonResponse
	TmallCarFinanceStatusSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarFinanceStatusSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarFinanceStatusSyncAPIResponseModel).Reset()
}

// TmallCarFinanceStatusSyncAPIResponseModel is 汽车金融状态同步 成功返回结果
type TmallCarFinanceStatusSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_finance_status_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 同步结果
	Data *CreditLoanStatusSyncResp `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarFinanceStatusSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Data = nil
}

var poolTmallCarFinanceStatusSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarFinanceStatusSyncAPIResponse)
	},
}

// GetTmallCarFinanceStatusSyncAPIResponse 从 sync.Pool 获取 TmallCarFinanceStatusSyncAPIResponse
func GetTmallCarFinanceStatusSyncAPIResponse() *TmallCarFinanceStatusSyncAPIResponse {
	return poolTmallCarFinanceStatusSyncAPIResponse.Get().(*TmallCarFinanceStatusSyncAPIResponse)
}

// ReleaseTmallCarFinanceStatusSyncAPIResponse 将 TmallCarFinanceStatusSyncAPIResponse 保存到 sync.Pool
func ReleaseTmallCarFinanceStatusSyncAPIResponse(v *TmallCarFinanceStatusSyncAPIResponse) {
	v.Reset()
	poolTmallCarFinanceStatusSyncAPIResponse.Put(v)
}
