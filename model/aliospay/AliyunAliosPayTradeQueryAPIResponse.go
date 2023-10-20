package aliospay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunAliosPayTradeQueryAPIResponse 查询支付结果接口 API返回值
// aliyun.alios.pay.trade.query
//
// 商户用来查询支付结果接口
type AliyunAliosPayTradeQueryAPIResponse struct {
	model.CommonResponse
	AliyunAliosPayTradeQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunAliosPayTradeQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunAliosPayTradeQueryAPIResponseModel).Reset()
}

// AliyunAliosPayTradeQueryAPIResponseModel is 查询支付结果接口 成功返回结果
type AliyunAliosPayTradeQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_alios_pay_trade_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	AliospayResponse *AliOSPayResponse `json:"aliospay_response,omitempty" xml:"aliospay_response,omitempty"`
}

// Reset 清空结构体
func (m *AliyunAliosPayTradeQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AliospayResponse = nil
}

var poolAliyunAliosPayTradeQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunAliosPayTradeQueryAPIResponse)
	},
}

// GetAliyunAliosPayTradeQueryAPIResponse 从 sync.Pool 获取 AliyunAliosPayTradeQueryAPIResponse
func GetAliyunAliosPayTradeQueryAPIResponse() *AliyunAliosPayTradeQueryAPIResponse {
	return poolAliyunAliosPayTradeQueryAPIResponse.Get().(*AliyunAliosPayTradeQueryAPIResponse)
}

// ReleaseAliyunAliosPayTradeQueryAPIResponse 将 AliyunAliosPayTradeQueryAPIResponse 保存到 sync.Pool
func ReleaseAliyunAliosPayTradeQueryAPIResponse(v *AliyunAliosPayTradeQueryAPIResponse) {
	v.Reset()
	poolAliyunAliosPayTradeQueryAPIResponse.Put(v)
}
