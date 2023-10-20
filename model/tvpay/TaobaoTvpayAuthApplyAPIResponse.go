package tvpay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayAuthApplyAPIResponse tv支付申请设备授权 API返回值
// taobao.tvpay.auth.apply
//
// 为用户在指定设备上申请支付授权
type TaobaoTvpayAuthApplyAPIResponse struct {
	model.CommonResponse
	TaobaoTvpayAuthApplyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTvpayAuthApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTvpayAuthApplyAPIResponseModel).Reset()
}

// TaobaoTvpayAuthApplyAPIResponseModel is tv支付申请设备授权 成功返回结果
type TaobaoTvpayAuthApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"tvpay_auth_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTvpayAuthApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTvpayAuthApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTvpayAuthApplyAPIResponse)
	},
}

// GetTaobaoTvpayAuthApplyAPIResponse 从 sync.Pool 获取 TaobaoTvpayAuthApplyAPIResponse
func GetTaobaoTvpayAuthApplyAPIResponse() *TaobaoTvpayAuthApplyAPIResponse {
	return poolTaobaoTvpayAuthApplyAPIResponse.Get().(*TaobaoTvpayAuthApplyAPIResponse)
}

// ReleaseTaobaoTvpayAuthApplyAPIResponse 将 TaobaoTvpayAuthApplyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTvpayAuthApplyAPIResponse(v *TaobaoTvpayAuthApplyAPIResponse) {
	v.Reset()
	poolTaobaoTvpayAuthApplyAPIResponse.Put(v)
}
