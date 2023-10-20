package alidoc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthNrRxPrescriptionGetAPIResponse 搜索处方详情 API返回值
// alibaba.alihealth.nr.rx.prescription.get
//
// 获取互联网医院处方详情
type AlibabaAlihealthNrRxPrescriptionGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthNrRxPrescriptionGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthNrRxPrescriptionGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthNrRxPrescriptionGetAPIResponseModel).Reset()
}

// AlibabaAlihealthNrRxPrescriptionGetAPIResponseModel is 搜索处方详情 成功返回结果
type AlibabaAlihealthNrRxPrescriptionGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_nr_rx_prescription_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 数据集
	DataList []PrescriptionSearchResultDto `json:"data_list,omitempty" xml:"data_list>prescription_search_result_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthNrRxPrescriptionGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DataList = m.DataList[:0]
}

var poolAlibabaAlihealthNrRxPrescriptionGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthNrRxPrescriptionGetAPIResponse)
	},
}

// GetAlibabaAlihealthNrRxPrescriptionGetAPIResponse 从 sync.Pool 获取 AlibabaAlihealthNrRxPrescriptionGetAPIResponse
func GetAlibabaAlihealthNrRxPrescriptionGetAPIResponse() *AlibabaAlihealthNrRxPrescriptionGetAPIResponse {
	return poolAlibabaAlihealthNrRxPrescriptionGetAPIResponse.Get().(*AlibabaAlihealthNrRxPrescriptionGetAPIResponse)
}

// ReleaseAlibabaAlihealthNrRxPrescriptionGetAPIResponse 将 AlibabaAlihealthNrRxPrescriptionGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthNrRxPrescriptionGetAPIResponse(v *AlibabaAlihealthNrRxPrescriptionGetAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthNrRxPrescriptionGetAPIResponse.Put(v)
}
