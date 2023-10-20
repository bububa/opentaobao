package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomVtDistributeQueryprotocolAPIResponse 通信业务外放协议查询 API返回值
// alibaba.alicom.vt.distribute.queryprotocol
//
// 通信业务外放协议查询
type AlibabaAlicomVtDistributeQueryprotocolAPIResponse struct {
	model.CommonResponse
	AlibabaAlicomVtDistributeQueryprotocolAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlicomVtDistributeQueryprotocolAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlicomVtDistributeQueryprotocolAPIResponseModel).Reset()
}

// AlibabaAlicomVtDistributeQueryprotocolAPIResponseModel is 通信业务外放协议查询 成功返回结果
type AlibabaAlicomVtDistributeQueryprotocolAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alicom_vt_distribute_queryprotocol_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlicomVtDistributeQueryprotocolAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlicomVtDistributeQueryprotocolAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlicomVtDistributeQueryprotocolAPIResponse)
	},
}

// GetAlibabaAlicomVtDistributeQueryprotocolAPIResponse 从 sync.Pool 获取 AlibabaAlicomVtDistributeQueryprotocolAPIResponse
func GetAlibabaAlicomVtDistributeQueryprotocolAPIResponse() *AlibabaAlicomVtDistributeQueryprotocolAPIResponse {
	return poolAlibabaAlicomVtDistributeQueryprotocolAPIResponse.Get().(*AlibabaAlicomVtDistributeQueryprotocolAPIResponse)
}

// ReleaseAlibabaAlicomVtDistributeQueryprotocolAPIResponse 将 AlibabaAlicomVtDistributeQueryprotocolAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlicomVtDistributeQueryprotocolAPIResponse(v *AlibabaAlicomVtDistributeQueryprotocolAPIResponse) {
	v.Reset()
	poolAlibabaAlicomVtDistributeQueryprotocolAPIResponse.Put(v)
}
