package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPhoneBankCreditProcessAPIResponse 虚拟话费任务银行信用卡办理进度回传 API返回值
// taobao.phone.bank.credit.process
//
// 虚拟话费任务银行信用卡办理进度回传
type TaobaoPhoneBankCreditProcessAPIResponse struct {
	model.CommonResponse
	TaobaoPhoneBankCreditProcessAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPhoneBankCreditProcessAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPhoneBankCreditProcessAPIResponseModel).Reset()
}

// TaobaoPhoneBankCreditProcessAPIResponseModel is 虚拟话费任务银行信用卡办理进度回传 成功返回结果
type TaobaoPhoneBankCreditProcessAPIResponseModel struct {
	XMLName xml.Name `xml:"phone_bank_credit_process_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的结果码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 响应说明
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 响应结果
	Data *BankCreditResponse `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPhoneBankCreditProcessAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizCode = ""
	m.Message = ""
	m.Data = nil
}

var poolTaobaoPhoneBankCreditProcessAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPhoneBankCreditProcessAPIResponse)
	},
}

// GetTaobaoPhoneBankCreditProcessAPIResponse 从 sync.Pool 获取 TaobaoPhoneBankCreditProcessAPIResponse
func GetTaobaoPhoneBankCreditProcessAPIResponse() *TaobaoPhoneBankCreditProcessAPIResponse {
	return poolTaobaoPhoneBankCreditProcessAPIResponse.Get().(*TaobaoPhoneBankCreditProcessAPIResponse)
}

// ReleaseTaobaoPhoneBankCreditProcessAPIResponse 将 TaobaoPhoneBankCreditProcessAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPhoneBankCreditProcessAPIResponse(v *TaobaoPhoneBankCreditProcessAPIResponse) {
	v.Reset()
	poolTaobaoPhoneBankCreditProcessAPIResponse.Put(v)
}
