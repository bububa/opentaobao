package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransPayRegisterCreateAPIResponse 提交支付服务开通 API返回值
// taobao.alitrip.axin.trans.pay.register.create
//
// 阿信供销平台-提交支付服务开通接口
type TaobaoAlitripAxinTransPayRegisterCreateAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripAxinTransPayRegisterCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransPayRegisterCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripAxinTransPayRegisterCreateAPIResponseModel).Reset()
}

// TaobaoAlitripAxinTransPayRegisterCreateAPIResponseModel is 提交支付服务开通 成功返回结果
type TaobaoAlitripAxinTransPayRegisterCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_axin_trans_pay_register_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripAxinTransPayRegisterCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransPayRegisterCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripAxinTransPayRegisterCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripAxinTransPayRegisterCreateAPIResponse)
	},
}

// GetTaobaoAlitripAxinTransPayRegisterCreateAPIResponse 从 sync.Pool 获取 TaobaoAlitripAxinTransPayRegisterCreateAPIResponse
func GetTaobaoAlitripAxinTransPayRegisterCreateAPIResponse() *TaobaoAlitripAxinTransPayRegisterCreateAPIResponse {
	return poolTaobaoAlitripAxinTransPayRegisterCreateAPIResponse.Get().(*TaobaoAlitripAxinTransPayRegisterCreateAPIResponse)
}

// ReleaseTaobaoAlitripAxinTransPayRegisterCreateAPIResponse 将 TaobaoAlitripAxinTransPayRegisterCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripAxinTransPayRegisterCreateAPIResponse(v *TaobaoAlitripAxinTransPayRegisterCreateAPIResponse) {
	v.Reset()
	poolTaobaoAlitripAxinTransPayRegisterCreateAPIResponse.Put(v)
}
