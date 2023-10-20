package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransRefundCreateAPIResponse 阿信创建退款单 API返回值
// taobao.alitrip.axin.trans.refund.create
//
// 阿信供销平台-创建退款单服务
type TaobaoAlitripAxinTransRefundCreateAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripAxinTransRefundCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransRefundCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripAxinTransRefundCreateAPIResponseModel).Reset()
}

// TaobaoAlitripAxinTransRefundCreateAPIResponseModel is 阿信创建退款单 成功返回结果
type TaobaoAlitripAxinTransRefundCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_axin_trans_refund_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripAxinTransRefundCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransRefundCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripAxinTransRefundCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripAxinTransRefundCreateAPIResponse)
	},
}

// GetTaobaoAlitripAxinTransRefundCreateAPIResponse 从 sync.Pool 获取 TaobaoAlitripAxinTransRefundCreateAPIResponse
func GetTaobaoAlitripAxinTransRefundCreateAPIResponse() *TaobaoAlitripAxinTransRefundCreateAPIResponse {
	return poolTaobaoAlitripAxinTransRefundCreateAPIResponse.Get().(*TaobaoAlitripAxinTransRefundCreateAPIResponse)
}

// ReleaseTaobaoAlitripAxinTransRefundCreateAPIResponse 将 TaobaoAlitripAxinTransRefundCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripAxinTransRefundCreateAPIResponse(v *TaobaoAlitripAxinTransRefundCreateAPIResponse) {
	v.Reset()
	poolTaobaoAlitripAxinTransRefundCreateAPIResponse.Put(v)
}
