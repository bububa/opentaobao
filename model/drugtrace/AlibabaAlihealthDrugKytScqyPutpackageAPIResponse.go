package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyPutpackageAPIResponse 码拼箱 API返回值
// alibaba.alihealth.drug.kyt.scqy.putpackage
//
// 码拼箱接口
type AlibabaAlihealthDrugKytScqyPutpackageAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytScqyPutpackageAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqyPutpackageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytScqyPutpackageAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytScqyPutpackageAPIResponseModel is 码拼箱 成功返回结果
type AlibabaAlihealthDrugKytScqyPutpackageAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_putpackage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugKytScqyPutpackageResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqyPutpackageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytScqyPutpackageAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqyPutpackageAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytScqyPutpackageAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytScqyPutpackageAPIResponse
func GetAlibabaAlihealthDrugKytScqyPutpackageAPIResponse() *AlibabaAlihealthDrugKytScqyPutpackageAPIResponse {
	return poolAlibabaAlihealthDrugKytScqyPutpackageAPIResponse.Get().(*AlibabaAlihealthDrugKytScqyPutpackageAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytScqyPutpackageAPIResponse 将 AlibabaAlihealthDrugKytScqyPutpackageAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytScqyPutpackageAPIResponse(v *AlibabaAlihealthDrugKytScqyPutpackageAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytScqyPutpackageAPIResponse.Put(v)
}
