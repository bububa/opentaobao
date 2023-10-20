package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse 码流向查询 API返回值
// alibaba.alihealth.drug.code.kyt.querycodeflow
//
// 追溯码流向查询
type AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponseModel is 码流向查询 成功返回结果
type AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_querycodeflow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugCodeKytQuerycodeflowResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse)
	},
}

// GetAlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse
func GetAlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse() *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse {
	return poolAlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse.Get().(*AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse)
}

// ReleaseAlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse 将 AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse(v *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse.Put(v)
}
