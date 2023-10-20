package eleenterpriseordernew

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse 退单和申诉 API返回值
// alibaba.ele.enterprise.ordernew.getrefundinfo
//
// 退单和申诉
type AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponseModel is 退单和申诉 成功返回结果
type AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_ordernew_getrefundinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应code
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 响应信息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
	// 返回值信息
	EnterpriseData *StandardOrderTrackingInfoDto `json:"enterprise_data,omitempty" xml:"enterprise_data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
	m.EnterpriseData = nil
}

var poolAlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse)
	},
}

// GetAlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse
func GetAlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse() *AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse {
	return poolAlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse.Get().(*AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse)
}

// ReleaseAlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse 将 AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse(v *AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse.Put(v)
}
