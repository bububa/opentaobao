package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomVtDistributeorderCreateAPIResponse 通信业务外放下单 API返回值
// alibaba.alicom.vt.distributeorder.create
//
// 通信业务外放下单接口
type AlibabaAlicomVtDistributeorderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaAlicomVtDistributeorderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlicomVtDistributeorderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlicomVtDistributeorderCreateAPIResponseModel).Reset()
}

// AlibabaAlicomVtDistributeorderCreateAPIResponseModel is 通信业务外放下单 成功返回结果
type AlibabaAlicomVtDistributeorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alicom_vt_distributeorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlicomVtDistributeorderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlicomVtDistributeorderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlicomVtDistributeorderCreateAPIResponse)
	},
}

// GetAlibabaAlicomVtDistributeorderCreateAPIResponse 从 sync.Pool 获取 AlibabaAlicomVtDistributeorderCreateAPIResponse
func GetAlibabaAlicomVtDistributeorderCreateAPIResponse() *AlibabaAlicomVtDistributeorderCreateAPIResponse {
	return poolAlibabaAlicomVtDistributeorderCreateAPIResponse.Get().(*AlibabaAlicomVtDistributeorderCreateAPIResponse)
}

// ReleaseAlibabaAlicomVtDistributeorderCreateAPIResponse 将 AlibabaAlicomVtDistributeorderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlicomVtDistributeorderCreateAPIResponse(v *AlibabaAlicomVtDistributeorderCreateAPIResponse) {
	v.Reset()
	poolAlibabaAlicomVtDistributeorderCreateAPIResponse.Put(v)
}
