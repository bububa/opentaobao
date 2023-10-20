package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangConsignorderGiftBindingAPIResponse 赠品绑赠计算占用 API返回值
// alibaba.dchain.aoxiang.consignorder.gift.binding
//
// 赠品绑赠计算占用
type AlibabaDchainAoxiangConsignorderGiftBindingAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangConsignorderGiftBindingAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangConsignorderGiftBindingAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangConsignorderGiftBindingAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangConsignorderGiftBindingAPIResponseModel is 赠品绑赠计算占用 成功返回结果
type AlibabaDchainAoxiangConsignorderGiftBindingAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_consignorder_gift_binding_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BindingConsignorderGiftResponse *BindingConsignOrderGiftRequest `json:"binding_consignorder_gift_response,omitempty" xml:"binding_consignorder_gift_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangConsignorderGiftBindingAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BindingConsignorderGiftResponse = nil
}

var poolAlibabaDchainAoxiangConsignorderGiftBindingAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangConsignorderGiftBindingAPIResponse)
	},
}

// GetAlibabaDchainAoxiangConsignorderGiftBindingAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangConsignorderGiftBindingAPIResponse
func GetAlibabaDchainAoxiangConsignorderGiftBindingAPIResponse() *AlibabaDchainAoxiangConsignorderGiftBindingAPIResponse {
	return poolAlibabaDchainAoxiangConsignorderGiftBindingAPIResponse.Get().(*AlibabaDchainAoxiangConsignorderGiftBindingAPIResponse)
}

// ReleaseAlibabaDchainAoxiangConsignorderGiftBindingAPIResponse 将 AlibabaDchainAoxiangConsignorderGiftBindingAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangConsignorderGiftBindingAPIResponse(v *AlibabaDchainAoxiangConsignorderGiftBindingAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangConsignorderGiftBindingAPIResponse.Put(v)
}
