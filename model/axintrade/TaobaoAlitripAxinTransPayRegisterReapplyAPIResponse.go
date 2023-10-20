package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransPayRegisterReapplyAPIResponse 阿信支付入驻重新申请 API返回值
// taobao.alitrip.axin.trans.pay.register.reapply
//
// 阿信支付入驻重新申请
// 用于支付平台驳回，商户提交时的业务场景
type TaobaoAlitripAxinTransPayRegisterReapplyAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripAxinTransPayRegisterReapplyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransPayRegisterReapplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripAxinTransPayRegisterReapplyAPIResponseModel).Reset()
}

// TaobaoAlitripAxinTransPayRegisterReapplyAPIResponseModel is 阿信支付入驻重新申请 成功返回结果
type TaobaoAlitripAxinTransPayRegisterReapplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_axin_trans_pay_register_reapply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *BaseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransPayRegisterReapplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripAxinTransPayRegisterReapplyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripAxinTransPayRegisterReapplyAPIResponse)
	},
}

// GetTaobaoAlitripAxinTransPayRegisterReapplyAPIResponse 从 sync.Pool 获取 TaobaoAlitripAxinTransPayRegisterReapplyAPIResponse
func GetTaobaoAlitripAxinTransPayRegisterReapplyAPIResponse() *TaobaoAlitripAxinTransPayRegisterReapplyAPIResponse {
	return poolTaobaoAlitripAxinTransPayRegisterReapplyAPIResponse.Get().(*TaobaoAlitripAxinTransPayRegisterReapplyAPIResponse)
}

// ReleaseTaobaoAlitripAxinTransPayRegisterReapplyAPIResponse 将 TaobaoAlitripAxinTransPayRegisterReapplyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripAxinTransPayRegisterReapplyAPIResponse(v *TaobaoAlitripAxinTransPayRegisterReapplyAPIResponse) {
	v.Reset()
	poolTaobaoAlitripAxinTransPayRegisterReapplyAPIResponse.Put(v)
}
