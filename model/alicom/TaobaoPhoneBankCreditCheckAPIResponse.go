package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPhoneBankCreditCheckAPIResponse 虚拟话费任务银行信用卡办理检查 API返回值
// taobao.phone.bank.credit.check
//
// 虚拟话费任务银行信用卡办理检查
type TaobaoPhoneBankCreditCheckAPIResponse struct {
	model.CommonResponse
	TaobaoPhoneBankCreditCheckAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPhoneBankCreditCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPhoneBankCreditCheckAPIResponseModel).Reset()
}

// TaobaoPhoneBankCreditCheckAPIResponseModel is 虚拟话费任务银行信用卡办理检查 成功返回结果
type TaobaoPhoneBankCreditCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"phone_bank_credit_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 响应结果
	Data *BankCreditCheckResponse `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPhoneBankCreditCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizCode = ""
	m.Message = ""
	m.Data = nil
}

var poolTaobaoPhoneBankCreditCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPhoneBankCreditCheckAPIResponse)
	},
}

// GetTaobaoPhoneBankCreditCheckAPIResponse 从 sync.Pool 获取 TaobaoPhoneBankCreditCheckAPIResponse
func GetTaobaoPhoneBankCreditCheckAPIResponse() *TaobaoPhoneBankCreditCheckAPIResponse {
	return poolTaobaoPhoneBankCreditCheckAPIResponse.Get().(*TaobaoPhoneBankCreditCheckAPIResponse)
}

// ReleaseTaobaoPhoneBankCreditCheckAPIResponse 将 TaobaoPhoneBankCreditCheckAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPhoneBankCreditCheckAPIResponse(v *TaobaoPhoneBankCreditCheckAPIResponse) {
	v.Reset()
	poolTaobaoPhoneBankCreditCheckAPIResponse.Put(v)
}
