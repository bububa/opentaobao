package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransPaySignCheckAPIResponse 阿信支付宝验签服务 API返回值
// taobao.alitrip.axin.trans.pay.sign.check
//
// 阿信支付宝验签服务
type TaobaoAlitripAxinTransPaySignCheckAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripAxinTransPaySignCheckAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransPaySignCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripAxinTransPaySignCheckAPIResponseModel).Reset()
}

// TaobaoAlitripAxinTransPaySignCheckAPIResponseModel is 阿信支付宝验签服务 成功返回结果
type TaobaoAlitripAxinTransPaySignCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_axin_trans_pay_sign_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *BaseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransPaySignCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripAxinTransPaySignCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripAxinTransPaySignCheckAPIResponse)
	},
}

// GetTaobaoAlitripAxinTransPaySignCheckAPIResponse 从 sync.Pool 获取 TaobaoAlitripAxinTransPaySignCheckAPIResponse
func GetTaobaoAlitripAxinTransPaySignCheckAPIResponse() *TaobaoAlitripAxinTransPaySignCheckAPIResponse {
	return poolTaobaoAlitripAxinTransPaySignCheckAPIResponse.Get().(*TaobaoAlitripAxinTransPaySignCheckAPIResponse)
}

// ReleaseTaobaoAlitripAxinTransPaySignCheckAPIResponse 将 TaobaoAlitripAxinTransPaySignCheckAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripAxinTransPaySignCheckAPIResponse(v *TaobaoAlitripAxinTransPaySignCheckAPIResponse) {
	v.Reset()
	poolTaobaoAlitripAxinTransPaySignCheckAPIResponse.Put(v)
}
