package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOcApContractsignedGetAPIResponse 用户是否签署支付宝代扣协议 API返回值
// taobao.oc.ap.contractsigned.get
//
// 用户是否签署支付宝代扣协议
type TaobaoOcApContractsignedGetAPIResponse struct {
	model.CommonResponse
	TaobaoOcApContractsignedGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOcApContractsignedGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOcApContractsignedGetAPIResponseModel).Reset()
}

// TaobaoOcApContractsignedGetAPIResponseModel is 用户是否签署支付宝代扣协议 成功返回结果
type TaobaoOcApContractsignedGetAPIResponseModel struct {
	XMLName xml.Name `xml:"oc_ap_contractsigned_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用出错描述信息
	ErrorDescription string `json:"error_description,omitempty" xml:"error_description,omitempty"`
	// 是否开通
	ContractSign bool `json:"contract_sign,omitempty" xml:"contract_sign,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOcApContractsignedGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorDescription = ""
	m.ContractSign = false
}

var poolTaobaoOcApContractsignedGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOcApContractsignedGetAPIResponse)
	},
}

// GetTaobaoOcApContractsignedGetAPIResponse 从 sync.Pool 获取 TaobaoOcApContractsignedGetAPIResponse
func GetTaobaoOcApContractsignedGetAPIResponse() *TaobaoOcApContractsignedGetAPIResponse {
	return poolTaobaoOcApContractsignedGetAPIResponse.Get().(*TaobaoOcApContractsignedGetAPIResponse)
}

// ReleaseTaobaoOcApContractsignedGetAPIResponse 将 TaobaoOcApContractsignedGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOcApContractsignedGetAPIResponse(v *TaobaoOcApContractsignedGetAPIResponse) {
	v.Reset()
	poolTaobaoOcApContractsignedGetAPIResponse.Put(v)
}
