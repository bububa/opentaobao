package aliospay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunAliosPayRefundAPIResponse 退款接口 API返回值
// aliyun.alios.pay.refund
//
// 商户用来发起退款的接口
type AliyunAliosPayRefundAPIResponse struct {
	model.CommonResponse
	AliyunAliosPayRefundAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunAliosPayRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunAliosPayRefundAPIResponseModel).Reset()
}

// AliyunAliosPayRefundAPIResponseModel is 退款接口 成功返回结果
type AliyunAliosPayRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_alios_pay_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	AliospayResponse *AliOSPayResponse `json:"aliospay_response,omitempty" xml:"aliospay_response,omitempty"`
}

// Reset 清空结构体
func (m *AliyunAliosPayRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AliospayResponse = nil
}

var poolAliyunAliosPayRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunAliosPayRefundAPIResponse)
	},
}

// GetAliyunAliosPayRefundAPIResponse 从 sync.Pool 获取 AliyunAliosPayRefundAPIResponse
func GetAliyunAliosPayRefundAPIResponse() *AliyunAliosPayRefundAPIResponse {
	return poolAliyunAliosPayRefundAPIResponse.Get().(*AliyunAliosPayRefundAPIResponse)
}

// ReleaseAliyunAliosPayRefundAPIResponse 将 AliyunAliosPayRefundAPIResponse 保存到 sync.Pool
func ReleaseAliyunAliosPayRefundAPIResponse(v *AliyunAliosPayRefundAPIResponse) {
	v.Reset()
	poolAliyunAliosPayRefundAPIResponse.Put(v)
}
