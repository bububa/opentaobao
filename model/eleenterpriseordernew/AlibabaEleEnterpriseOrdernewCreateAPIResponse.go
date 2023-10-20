package eleenterpriseordernew

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseOrdernewCreateAPIResponse 创建订单 API返回值
// alibaba.ele.enterprise.ordernew.create
//
// 创建订单
type AlibabaEleEnterpriseOrdernewCreateAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseOrdernewCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseOrdernewCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseOrdernewCreateAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseOrdernewCreateAPIResponseModel is 创建订单 成功返回结果
type AlibabaEleEnterpriseOrdernewCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_ordernew_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应code
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 响应信息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
	// 结果
	EnterpriseData *EnterpriseData `json:"enterprise_data,omitempty" xml:"enterprise_data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseOrdernewCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
	m.EnterpriseData = nil
}

var poolAlibabaEleEnterpriseOrdernewCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseOrdernewCreateAPIResponse)
	},
}

// GetAlibabaEleEnterpriseOrdernewCreateAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseOrdernewCreateAPIResponse
func GetAlibabaEleEnterpriseOrdernewCreateAPIResponse() *AlibabaEleEnterpriseOrdernewCreateAPIResponse {
	return poolAlibabaEleEnterpriseOrdernewCreateAPIResponse.Get().(*AlibabaEleEnterpriseOrdernewCreateAPIResponse)
}

// ReleaseAlibabaEleEnterpriseOrdernewCreateAPIResponse 将 AlibabaEleEnterpriseOrdernewCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseOrdernewCreateAPIResponse(v *AlibabaEleEnterpriseOrdernewCreateAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseOrdernewCreateAPIResponse.Put(v)
}
