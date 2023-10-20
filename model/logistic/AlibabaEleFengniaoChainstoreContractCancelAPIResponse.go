package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoChainstoreContractCancelAPIResponse 门店解约接口 API返回值
// alibaba.ele.fengniao.chainstore.contract.cancel
//
// 调用成功后，门店和蜂鸟解除物流合同，不能再使用此门店推单
type AlibabaEleFengniaoChainstoreContractCancelAPIResponse struct {
	model.CommonResponse
	AlibabaEleFengniaoChainstoreContractCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoChainstoreContractCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleFengniaoChainstoreContractCancelAPIResponseModel).Reset()
}

// AlibabaEleFengniaoChainstoreContractCancelAPIResponseModel is 门店解约接口 成功返回结果
type AlibabaEleFengniaoChainstoreContractCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_fengniao_chainstore_contract_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msg
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoChainstoreContractCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
}

var poolAlibabaEleFengniaoChainstoreContractCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleFengniaoChainstoreContractCancelAPIResponse)
	},
}

// GetAlibabaEleFengniaoChainstoreContractCancelAPIResponse 从 sync.Pool 获取 AlibabaEleFengniaoChainstoreContractCancelAPIResponse
func GetAlibabaEleFengniaoChainstoreContractCancelAPIResponse() *AlibabaEleFengniaoChainstoreContractCancelAPIResponse {
	return poolAlibabaEleFengniaoChainstoreContractCancelAPIResponse.Get().(*AlibabaEleFengniaoChainstoreContractCancelAPIResponse)
}

// ReleaseAlibabaEleFengniaoChainstoreContractCancelAPIResponse 将 AlibabaEleFengniaoChainstoreContractCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleFengniaoChainstoreContractCancelAPIResponse(v *AlibabaEleFengniaoChainstoreContractCancelAPIResponse) {
	v.Reset()
	poolAlibabaEleFengniaoChainstoreContractCancelAPIResponse.Put(v)
}
