package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaophonebankcreditcheckAPIResponse 虚拟话费任务银行信用卡办理检查 API返回值
// taobao.phone.bank.credit.check
//
// 虚拟话费任务银行信用卡办理检查
type TaobaophonebankcreditcheckAPIResponse struct {
	model.CommonResponse
	TaobaophonebankcreditcheckAPIResponseModel
}

// TaobaophonebankcreditcheckAPIResponseModel is 虚拟话费任务银行信用卡办理检查 成功返回结果
type TaobaophonebankcreditcheckAPIResponseModel struct {
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
