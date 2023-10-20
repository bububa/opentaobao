package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransPayRegisterAuditAPIResponse 阿信支付入驻审核通知 API返回值
// taobao.alitrip.axin.trans.pay.register.audit
//
// 阿信支付入驻审核通知
type TaobaoAlitripAxinTransPayRegisterAuditAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripAxinTransPayRegisterAuditAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransPayRegisterAuditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripAxinTransPayRegisterAuditAPIResponseModel).Reset()
}

// TaobaoAlitripAxinTransPayRegisterAuditAPIResponseModel is 阿信支付入驻审核通知 成功返回结果
type TaobaoAlitripAxinTransPayRegisterAuditAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_axin_trans_pay_register_audit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransPayRegisterAuditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripAxinTransPayRegisterAuditAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripAxinTransPayRegisterAuditAPIResponse)
	},
}

// GetTaobaoAlitripAxinTransPayRegisterAuditAPIResponse 从 sync.Pool 获取 TaobaoAlitripAxinTransPayRegisterAuditAPIResponse
func GetTaobaoAlitripAxinTransPayRegisterAuditAPIResponse() *TaobaoAlitripAxinTransPayRegisterAuditAPIResponse {
	return poolTaobaoAlitripAxinTransPayRegisterAuditAPIResponse.Get().(*TaobaoAlitripAxinTransPayRegisterAuditAPIResponse)
}

// ReleaseTaobaoAlitripAxinTransPayRegisterAuditAPIResponse 将 TaobaoAlitripAxinTransPayRegisterAuditAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripAxinTransPayRegisterAuditAPIResponse(v *TaobaoAlitripAxinTransPayRegisterAuditAPIResponse) {
	v.Reset()
	poolTaobaoAlitripAxinTransPayRegisterAuditAPIResponse.Put(v)
}
