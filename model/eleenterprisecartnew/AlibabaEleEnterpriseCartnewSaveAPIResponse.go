package eleenterprisecartnew

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseCartnewSaveAPIResponse 新版创建购物车 API返回值
// alibaba.ele.enterprise.cartnew.save
//
// 新版创建购物车
type AlibabaEleEnterpriseCartnewSaveAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseCartnewSaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseCartnewSaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseCartnewSaveAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseCartnewSaveAPIResponseModel is 新版创建购物车 成功返回结果
type AlibabaEleEnterpriseCartnewSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_cartnew_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 状态消息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
	// 系统自动生成
	EnterpriseDatas *EnterpriseData `json:"enterprise_datas,omitempty" xml:"enterprise_datas,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseCartnewSaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
	m.EnterpriseDatas = nil
}

var poolAlibabaEleEnterpriseCartnewSaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseCartnewSaveAPIResponse)
	},
}

// GetAlibabaEleEnterpriseCartnewSaveAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseCartnewSaveAPIResponse
func GetAlibabaEleEnterpriseCartnewSaveAPIResponse() *AlibabaEleEnterpriseCartnewSaveAPIResponse {
	return poolAlibabaEleEnterpriseCartnewSaveAPIResponse.Get().(*AlibabaEleEnterpriseCartnewSaveAPIResponse)
}

// ReleaseAlibabaEleEnterpriseCartnewSaveAPIResponse 将 AlibabaEleEnterpriseCartnewSaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseCartnewSaveAPIResponse(v *AlibabaEleEnterpriseCartnewSaveAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseCartnewSaveAPIResponse.Put(v)
}
