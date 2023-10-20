package tbrefund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRpRefundInterceptAPIResponse 卖家发起拦截 API返回值
// taobao.rp.refund.intercept
//
// 卖家发起拦截
type TaobaoRpRefundInterceptAPIResponse struct {
	model.CommonResponse
	TaobaoRpRefundInterceptAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRpRefundInterceptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRpRefundInterceptAPIResponseModel).Reset()
}

// TaobaoRpRefundInterceptAPIResponseModel is 卖家发起拦截 成功返回结果
type TaobaoRpRefundInterceptAPIResponseModel struct {
	XMLName xml.Name `xml:"rp_refund_intercept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRpRefundInterceptAPIResponseModel) Reset() {
	m.RequestId = ""
}

var poolTaobaoRpRefundInterceptAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRpRefundInterceptAPIResponse)
	},
}

// GetTaobaoRpRefundInterceptAPIResponse 从 sync.Pool 获取 TaobaoRpRefundInterceptAPIResponse
func GetTaobaoRpRefundInterceptAPIResponse() *TaobaoRpRefundInterceptAPIResponse {
	return poolTaobaoRpRefundInterceptAPIResponse.Get().(*TaobaoRpRefundInterceptAPIResponse)
}

// ReleaseTaobaoRpRefundInterceptAPIResponse 将 TaobaoRpRefundInterceptAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRpRefundInterceptAPIResponse(v *TaobaoRpRefundInterceptAPIResponse) {
	v.Reset()
	poolTaobaoRpRefundInterceptAPIResponse.Put(v)
}
