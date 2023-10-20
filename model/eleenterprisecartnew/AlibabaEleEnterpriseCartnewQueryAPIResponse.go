package eleenterprisecartnew

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseCartnewQueryAPIResponse 新版购物车查询 API返回值
// alibaba.ele.enterprise.cartnew.query
//
// 新版购物车查询
type AlibabaEleEnterpriseCartnewQueryAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseCartnewQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseCartnewQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseCartnewQueryAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseCartnewQueryAPIResponseModel is 新版购物车查询 成功返回结果
type AlibabaEleEnterpriseCartnewQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_cartnew_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回码
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 返回信息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
	// 系统自动生成
	EnterpriseData *EnterpriseData `json:"enterprise_data,omitempty" xml:"enterprise_data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseCartnewQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
	m.EnterpriseData = nil
}

var poolAlibabaEleEnterpriseCartnewQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseCartnewQueryAPIResponse)
	},
}

// GetAlibabaEleEnterpriseCartnewQueryAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseCartnewQueryAPIResponse
func GetAlibabaEleEnterpriseCartnewQueryAPIResponse() *AlibabaEleEnterpriseCartnewQueryAPIResponse {
	return poolAlibabaEleEnterpriseCartnewQueryAPIResponse.Get().(*AlibabaEleEnterpriseCartnewQueryAPIResponse)
}

// ReleaseAlibabaEleEnterpriseCartnewQueryAPIResponse 将 AlibabaEleEnterpriseCartnewQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseCartnewQueryAPIResponse(v *AlibabaEleEnterpriseCartnewQueryAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseCartnewQueryAPIResponse.Put(v)
}
