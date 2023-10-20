package alidoc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthAlidocDrugStoreAddAPIResponse gsk新增药店 API返回值
// alibaba.alihealth.alidoc.drug.store.add
//
// GSK上传药店信息
type AlibabaAlihealthAlidocDrugStoreAddAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthAlidocDrugStoreAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthAlidocDrugStoreAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthAlidocDrugStoreAddAPIResponseModel).Reset()
}

// AlibabaAlihealthAlidocDrugStoreAddAPIResponseModel is gsk新增药店 成功返回结果
type AlibabaAlihealthAlidocDrugStoreAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_alidoc_drug_store_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errCode
	ErrorKode string `json:"error_kode,omitempty" xml:"error_kode,omitempty"`
	// errMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// success
	Successed bool `json:"successed,omitempty" xml:"successed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthAlidocDrugStoreAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorKode = ""
	m.ErrorMessage = ""
	m.Successed = false
}

var poolAlibabaAlihealthAlidocDrugStoreAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthAlidocDrugStoreAddAPIResponse)
	},
}

// GetAlibabaAlihealthAlidocDrugStoreAddAPIResponse 从 sync.Pool 获取 AlibabaAlihealthAlidocDrugStoreAddAPIResponse
func GetAlibabaAlihealthAlidocDrugStoreAddAPIResponse() *AlibabaAlihealthAlidocDrugStoreAddAPIResponse {
	return poolAlibabaAlihealthAlidocDrugStoreAddAPIResponse.Get().(*AlibabaAlihealthAlidocDrugStoreAddAPIResponse)
}

// ReleaseAlibabaAlihealthAlidocDrugStoreAddAPIResponse 将 AlibabaAlihealthAlidocDrugStoreAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthAlidocDrugStoreAddAPIResponse(v *AlibabaAlihealthAlidocDrugStoreAddAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthAlidocDrugStoreAddAPIResponse.Put(v)
}
