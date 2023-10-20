package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangReceiverinfoQueryAPIResponse 供应链优仓即时配隐私小号查询 API返回值
// alibaba.dchain.aoxiang.receiverinfo.query
//
// 供应链优仓即时配隐私小号查询
type AlibabaDchainAoxiangReceiverinfoQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangReceiverinfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangReceiverinfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangReceiverinfoQueryAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangReceiverinfoQueryAPIResponseModel is 供应链优仓即时配隐私小号查询 成功返回结果
type AlibabaDchainAoxiangReceiverinfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_receiverinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户信息
	OrderReceiverPrivacyResponse *OrderReceiverPrivacyResponse `json:"order_receiver_privacy_response,omitempty" xml:"order_receiver_privacy_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangReceiverinfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderReceiverPrivacyResponse = nil
}

var poolAlibabaDchainAoxiangReceiverinfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangReceiverinfoQueryAPIResponse)
	},
}

// GetAlibabaDchainAoxiangReceiverinfoQueryAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangReceiverinfoQueryAPIResponse
func GetAlibabaDchainAoxiangReceiverinfoQueryAPIResponse() *AlibabaDchainAoxiangReceiverinfoQueryAPIResponse {
	return poolAlibabaDchainAoxiangReceiverinfoQueryAPIResponse.Get().(*AlibabaDchainAoxiangReceiverinfoQueryAPIResponse)
}

// ReleaseAlibabaDchainAoxiangReceiverinfoQueryAPIResponse 将 AlibabaDchainAoxiangReceiverinfoQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangReceiverinfoQueryAPIResponse(v *AlibabaDchainAoxiangReceiverinfoQueryAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangReceiverinfoQueryAPIResponse.Put(v)
}
